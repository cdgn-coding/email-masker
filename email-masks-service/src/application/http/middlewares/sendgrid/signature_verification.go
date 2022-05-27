package sendgrid

import (
	"net/http"
)

type signatureVerificationHandler struct {
	next http.Handler
}

func (s signatureVerificationHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// TODO perform signature verification
	s.next.ServeHTTP(writer, request)
}

func NewSignatureVerificationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return signatureVerificationHandler{next: next}
	}
}
