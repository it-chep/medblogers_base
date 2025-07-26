package interceptor

import (
	"context"

	"go.uber.org/ratelimit"
	"google.golang.org/grpc"
)

// todo вынести в конфиг
const maxRPS = 200

var limiter = ratelimit.New(maxRPS)

func RateLimitInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	limiter.Take()
	return handler(ctx, req)
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
