package interceptor

import (
	"context"
	"medblogers_base/internal/pkg/logger"

	"google.golang.org/grpc"
)

// LoggerInterceptor добавляет логгер в контекст gRPC вызова
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log := logger.New()
	ctx = logger.ContextWithLogger(ctx, log)
	return handler(ctx, req)
}
