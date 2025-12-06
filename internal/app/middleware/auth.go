package middleware

import (
	pkgctx "medblogers_base/internal/pkg/context"
	"medblogers_base/internal/pkg/token"
	"net/http"
)

type JWTConfig interface {
	GetJWTSecret() string
}

func EmailMiddleware(cfg JWTConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			claims, _ := token.AccessClaimsFromRequest(authHeader, cfg.GetJWTSecret())
			if claims == nil {
				next.ServeHTTP(w, r)
				return
			}

			r = r.WithContext(pkgctx.WithEmailContext(r.Context(), claims.Email))
			next.ServeHTTP(w, r)
		})
	}
}
