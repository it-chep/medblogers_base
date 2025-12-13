package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	"medblogers_base/internal/pkg/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata not found")
	}

	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		return nil, status.Error(codes.PermissionDenied, "Not Claims In Request")
	}

	claims, err := token.AccessClaimsFromRequest(authHeader[0], i.config.JWTConfig.Secret)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "err claims")
	}
	if claims == nil {
		return nil, status.Error(codes.PermissionDenied, "Not Claims In Request")
	}

	return &desc.CheckResponse{}, nil
}
