package auth

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	"medblogers_base/internal/pkg/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	claims, err := token.RefreshClaimsFromContext(ctx, i.config.JWTConfig.RefreshSecret)
	if err != nil {
		return nil, err
	}
	if claims == nil {
		return nil, status.Error(codes.Internal, "Not Claims In Request")
	}

	return &desc.CheckResponse{}, nil
}
