package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"medblogers_base/internal/pkg/logger"
	"time"
)

// ResponseTimeInterceptor логирует время выполнения запроса и статус ответа
func ResponseTimeInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	res, err := handler(ctx, req)
	end := time.Since(start)
	logger.Message(ctx, fmt.Sprintf("Запрос выполнился за: %v", end))
	return res, err
}
