package middleware

import (
	"fmt"
	"medblogers_base/internal/pkg/logger"
	"net/http"
	"time"
)

// ResponseTimeMiddleware логирует время выполнения запроса и статус ответа
func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start := time.Now()
		next.ServeHTTP(w, r)
		//end := time.Since(start)
		//logger.Message(r.Context(), fmt.Sprintf("Запрос выполнился за: %v", end))
		// todo залогать отдать метрику
	})
}
