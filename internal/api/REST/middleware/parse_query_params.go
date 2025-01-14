package middleware

import (
	"context"
	"net/http"
	"net/url"
)

// ParseQueryParams parses query parameters into a DTO and stores it in the context
func ParseQueryParams[T any](contextKey string, parseFunc func(url.Values) (T, error), next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dto, err := parseFunc(r.URL.Query())
		if err != nil {
			http.Error(w, "Invalid query parameters", http.StatusBadRequest)
			return
		}

		// Attach the DTO to the request context using the custom key
		ctx := context.WithValue(r.Context(), contextKey, dto)

		// Call the next handler with the modified context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
