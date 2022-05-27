package sendgrid

import (
	"email-masks-service/src/business/entities"
	"email-masks-service/src/business/usecases"
	"email-masks-service/src/infrastructure/configuration"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sendgrid/sendgrid-go/helpers/inbound"
	"net/http"
)

type InboundEmailController struct {
	redirectEmail usecases.CommandUseCase[*entities.Email]
	logger        configuration.Logger
}

func NewInboundEmailController(
	redirectEmail usecases.CommandUseCase[*entities.Email],
	logger configuration.Logger) *InboundEmailController {
	return &InboundEmailController{redirectEmail: redirectEmail, logger: logger}
}

var invalidInboundEmailRequest = errors.New("http request is not a valid inbound email")

var inboundEmailRequestNotSecure = errors.New("inbound email not passed security assessments")

var tooMuchRecipients = errors.New("inbound email has too much recipients")

func (controller InboundEmailController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	email, err := controller.parseInboundEmail(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(email)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.redirectEmail.Execute(email)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller InboundEmailController) parseInboundEmail(request *http.Request) (*entities.Email, error) {
	inboundEmail, err := inbound.Parse(request)

	if err != nil {
		return nil, fmt.Errorf("invalid request. %v. %w", err, invalidInboundEmailRequest)
	}

	err = inboundEmail.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid request. %v. %w", err, inboundEmailRequestNotSecure)
	}

	if len(inboundEmail.Envelope.To) > 1 {
		return nil, fmt.Errorf("invalid request. %v. %w", err, tooMuchRecipients)
	}

	email := &entities.Email{
		From:    inboundEmail.Envelope.From,
		To:      inboundEmail.Envelope.To[0],
		Subject: inboundEmail.Headers["Subject"],
		Content: inboundEmail.TextBody,
		HTML:    inboundEmail.Body["text/html"],
	}

	return email, nil
}
