package http

import (
	"email-masks-service/src/business/entities"
	"email-masks-service/src/business/usecases"
	"email-masks-service/src/infrastructure/configuration"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type InboundEmailController struct {
	redirectEmail usecases.CommandUseCase[entities.Email]
	logger        configuration.Logger
}

func (controller InboundEmailController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// todo read from request
	inboundEmail := entities.Email{
		From:    "",
		To:      "",
		Subject: "",
		Content: "",
		HTML:    "",
	}

	validate := validator.New()
	err := validate.Struct(inboundEmail)
	if err != nil {
		controller.logger.Warn("Inbound email with unexpected values. Received %+v. %v", inboundEmail, err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.redirectEmail.Execute(inboundEmail)

	if err != nil {
		controller.logger.Warn("Error redirecting email. %v", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
