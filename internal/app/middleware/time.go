package middleware

import (
	"net/http"
	"time"
)

// ResponseTimeMiddleware логирует время выполнения запроса и статус ответа
func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		_ = time.Since(start)
		// todo залогать отдать метрику
	})
}
