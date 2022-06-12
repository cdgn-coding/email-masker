package masks

import (
	masksUseCases "email-masks-service/src/business/usecases/masks"
	"email-masks-service/src/infrastructure/configuration"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type CreateEmailMaskController struct {
	createMask masksUseCases.ICreateMaskUseCase
	logger     configuration.Logger
}

func NewCreateEmailMaskController(createMask masksUseCases.ICreateMaskUseCase, logger configuration.Logger) *CreateEmailMaskController {
	return &CreateEmailMaskController{createMask: createMask, logger: logger}
}

type createEmailMaskBody struct {
	Name        string `validate:"required" json:"name"`
	Description string `json:"description"`
}

var InvalidRequestBody = errors.New("error while decoding request body")
var CreateMaskError = errors.New("error creating mask")

func (c CreateEmailMaskController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body createEmailMaskBody
	var err error

	err = json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		msg := fmt.Errorf("%w. %v", InvalidRequestBody, err).Error()
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		msg := fmt.Errorf("%w. %v", InvalidRequestBody, err).Error()
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	params := mux.Vars(request)
	userID := params["userID"]

	err = c.createMask.Execute(&masksUseCases.CreateMaskInput{
		Name:        body.Name,
		Description: body.Description,
		UserID:      userID,
	})
	if err != nil {
		msg := fmt.Errorf("%w. %v", CreateMaskError, err).Error()
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
