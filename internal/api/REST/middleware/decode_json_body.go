package middleware

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// DecodeJSONBody is a middleware that decodes the request body into a specified DTO
func DecodeJSONBody[T any](contextKey string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dto T
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read request body", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &dto)
		if err != nil {
			http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
			return
		}

		// Attach the DTO to the request context using the custom key
		ctx := context.WithValue(r.Context(), contextKey, dto)

		// Call the next handler with the modified context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
