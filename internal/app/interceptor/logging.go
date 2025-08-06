package interceptor

import (
	"context"
	"medblogers_base/internal/pkg/logger"

	"google.golang.org/grpc"
)

// LoggerInterceptor добавляет логгер в контекст gRPC вызова
func LoggerInterceptor(l *logger.Logger) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(logger.ContextWithLogger(ctx, l), req)
	}
}
