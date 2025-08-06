package middleware

import (
	"medblogers_base/internal/pkg/logger"
	"net/http"
)

// LoggerMiddleware добавляет логгер в контекст запроса
func LoggerMiddleware(l *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = logger.ContextWithLogger(ctx, l)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
