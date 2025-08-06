package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var allowedOrigins = map[string]bool{
	"https://doctors.readyschool.ru": true,
	"http://localhost:3000":          true, // для разработки
	"http://localhost:8080":          true, // для разработки
	"http://127.0.0.1:8080":          true, // для разработки
	"http://0.0.0.0:8080":            true, // для разработки
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Проверка origin через метаданные
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		origins := md.Get("origin")
		if len(origins) > 0 && !allowedOrigins[origins[0]] {
			return nil, status.Error(codes.PermissionDenied, "origin not allowed")
		}
	}

	return handler(ctx, req)
}
