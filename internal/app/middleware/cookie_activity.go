package middleware

import (
	"context"
	"net/http"
)

type cookieActivityUpdater interface {
	Do(ctx context.Context, cookieID string) error
}

// CookieActivityMiddleware обновляет время активности cookie пользователя.
func CookieActivityMiddleware(action cookieActivityUpdater) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookieID := r.Header.Get("cookie_id")
			if cookieID == "" {
				cookieID = r.Header.Get("Cookie-Id")
			}
			if cookieID == "" {
				next.ServeHTTP(w, r)
				return
			}

			_ = action.Do(r.Context(), cookieID)
			next.ServeHTTP(w, r)
		})
	}
}
