package http

import (
	"email-masks-service/src/business/usecases/emails"
	"net/http"
)

type InboundEmailController struct {
	redirectEmail *emails.RedirectEmailUseCase
}

func (i InboundEmailController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// TODO
}
