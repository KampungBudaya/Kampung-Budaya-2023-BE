package middleware

import (
	"net/http"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
)

func Guest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		if _, exists := headers["Authorization"]; exists {
			response.Fail(w, http.StatusUnauthorized, "Authorization Header Exists On Guest Only Endpoint")
			return
		}

		next.ServeHTTP(w, r)
	})
}
