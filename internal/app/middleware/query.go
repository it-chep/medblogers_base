package middleware

import (
	pkgctx "medblogers_base/internal/pkg/context"
	"net/http"
)

// ParseQueryMiddleware returns http middleware which parse query params
func ParseQueryMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			values := request.URL.Query()
			next.ServeHTTP(writer, request.WithContext(pkgctx.WithValues(request.Context(), values)))
		})
	}
}
