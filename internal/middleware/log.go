package middleware

import (
	"context"
	"log/slog"
	"net/http"
)

func AddLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "logger", logger))
			next.ServeHTTP(w, r)
		})
	}
}
