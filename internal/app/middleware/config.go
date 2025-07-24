package middleware

import (
	pkgConfig "medblogers_base/internal/pkg/config"
	"net/http"
)

func ConfigMiddleware(config pkgConfig.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := pkgConfig.ContextWithConfig(r.Context(), config)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
