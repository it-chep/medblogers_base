package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"medblogers_base/internal/pkg/logger"
)

// LoggerInterceptor добавляет логгер в контекст gRPC вызова
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log := logger.New()
	ctx = logger.ContextWithLogger(ctx, log)
	return handler(ctx, req)
}
