package middleware

import (
	"net/http"

	"go.uber.org/ratelimit"
)

// todo вынести в конфиг
const maxRPS = 200

var limiter = ratelimit.New(maxRPS)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter.Take()
		next.ServeHTTP(w, r)
	})
}

// Базовая реализация лимитера, пока не понятно какой буду пользоваться

//var (
//	logger  = zap.NewExample()
//	limiter = rate.NewLimiter(100, 10) // 100 RPS, burst=10
//)
//
//func RateLimitMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if !limiter.Allow() {
//			// Логируем факт превышения лимита
//			logger.Warn("Rate limit exceeded",
//				zap.String("ip", r.RemoteAddr),
//				zap.String("path", r.URL.Path),
//				zap.Time("time", time.Now()),
//			)
//
//			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
//			return
//		}
//		next.ServeHTTP(w, r)
//	})
//}
