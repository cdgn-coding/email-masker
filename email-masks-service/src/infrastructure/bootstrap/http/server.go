package http

import (
	sendgridControllers "email-masks-service/src/application/http/controllers/sendgrid"
	"email-masks-service/src/business/usecases/emails"
	"email-masks-service/src/infrastructure/configuration"
	"email-masks-service/src/infrastructure/drivers/auth0"
	"email-masks-service/src/infrastructure/drivers/postgresql"
	sendgridServices "email-masks-service/src/infrastructure/drivers/sendgrid"
	"github.com/sendgrid/sendgrid-go"
	"net/http"
)

type Server struct {
	sendgridEmailController http.Handler
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

	redirectEmailUseCase := emails.NewRedirectEmailUseCase(outboundEmailService, maskMappingService, usersService)
	sendgridEmailController := sendgridControllers.NewInboundEmailController(redirectEmailUseCase, logger)

	return &Server{
		sendgridEmailController,
	}
}

func (s Server) Run() {

}
