package http

import (
	"email-masks-service/src/application/http/controllers/masks"
	sendgridControllers "email-masks-service/src/application/http/controllers/sendgrid"
	"email-masks-service/src/application/http/middlewares/auth"
	sendgridMiddlewares "email-masks-service/src/application/http/middlewares/sendgrid"
	emailUseCases "email-masks-service/src/business/usecases/emails"
	masks2 "email-masks-service/src/business/usecases/masks"
	"email-masks-service/src/infrastructure/configuration"
	"email-masks-service/src/infrastructure/drivers/auth0"
	"email-masks-service/src/infrastructure/drivers/postgresql"
	sendgridServices "email-masks-service/src/infrastructure/drivers/sendgrid"
	"github.com/gorilla/mux"
	"github.com/sendgrid/sendgrid-go"
	"net/http"
)

type Server struct {
	config                        configuration.Config
	logger                        configuration.Logger
	sendgridEmailController       http.Handler
	createMaskController          http.Handler
	sendgridSignatureVerification mux.MiddlewareFunc
	jwtAuthorization              mux.MiddlewareFunc
}

func NewServer() *Server {
	config := configuration.LoadConfig()
	loggerLevel := config.GetString("logger.level")
	logger := configuration.NewLogger(loggerLevel)

	sendgridKey := config.GetString("sendgrid.clientKey")
	sendgridClient := sendgrid.NewSendClient(sendgridKey)
	outboundEmailService := sendgridServices.NewOutboundEmailService(sendgridClient)

	postgresConnection := postgresql.NewConnection(config.GetString("postgres.dsn"))
	maskMappingService := postgresql.NewPostgresMaskMappingService(postgresConnection)
	managementClient := auth0.NewManagementClient(
		config.GetString("auth0.domain"),
		config.GetString("auth0.clientID"),
		config.GetString("auth0.clientSecret"))
	usersService := auth0.NewAuth0UsersService(managementClient)

	redirectEmailUseCase := emailUseCases.NewRedirectEmailUseCase(outboundEmailService, maskMappingService, usersService)
	sendgridEmailController := sendgridControllers.NewInboundEmailController(redirectEmailUseCase, logger)
	sendgridSignatureVerification := sendgridMiddlewares.NewSignatureVerificationMiddleware()

	maskDomain := config.GetString("emails.domain")
	createMaskUseCase := masks2.NewCreateMaskUseCase(maskMappingService, usersService, maskDomain)
	createMaskController := masks.NewCreateEmailMaskController(createMaskUseCase, logger)

	jwtAuthorization := auth.NewAuthVerificationHandler()

	return &Server{
		config:                        config,
		logger:                        logger,
		sendgridEmailController:       sendgridEmailController,
		sendgridSignatureVerification: sendgridSignatureVerification,
		createMaskController:          createMaskController,
		jwtAuthorization:              jwtAuthorization,
	}
}

func (s Server) Run() {
	s.logger.Info("Starting up server", s.config)
	router := mux.NewRouter()
	sendgridWebhooksRouter := router.PathPrefix("/sendgrid").Subrouter()
	sendgridWebhooksRouter.Use(s.sendgridSignatureVerification)
	sendgridWebhooksRouter.Handle("/email", s.sendgridEmailController).Methods(http.MethodPost)

	webservice := router.PathPrefix("/users/{userID}/masks/email").Subrouter()
	webservice.Use(s.jwtAuthorization)
	webservice.Handle("/", s.createMaskController).Methods(http.MethodPost)

	s.logger.Fatal(http.ListenAndServe(s.config.GetString("http.port"), router))
}
