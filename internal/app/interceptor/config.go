package interceptor

import (
	"context"
	pkgConfig "medblogers_base/internal/pkg/config"

	"google.golang.org/grpc"
)

func ConfigInterceptor(config pkgConfig.Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = pkgConfig.ContextWithConfig(ctx, config)
		return handler(ctx, req)
	}
}
