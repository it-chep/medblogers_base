package interceptor

import (
	"context"
	"google.golang.org/grpc"
	pkgConfig "medblogers_base/internal/pkg/config"
)

func ConfigInterceptor(config pkgConfig.Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = pkgConfig.ContextWithConfig(ctx, config)
		return handler(ctx, req)
	}
}
