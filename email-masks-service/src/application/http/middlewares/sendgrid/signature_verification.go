package sendgrid

import "net/http"

func SignatureVerification(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// todo perform signature verification
		next.ServeHTTP(writer, request)
	}
}
