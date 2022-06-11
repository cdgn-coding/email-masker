package auth

import (
	"net/http"
)

type authVerificationHandler struct {
	next http.Handler
}

func (s authVerificationHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// TODO perform signature verification
	s.next.ServeHTTP(writer, request)
}

func NewAuthVerificationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return authVerificationHandler{next: next}
	}
}
