package auth

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	"medblogers_base/internal/pkg/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Refresh(ctx context.Context, req *desc.RefreshRequest) (*desc.RefreshResponse, error) {
	claims, err := token.RefreshClaimsFromContext(ctx, "")
	if err != nil {
		return nil, err
	}
	if claims == nil {
		return nil, status.Error(codes.Internal, "Not Claims In Request")
	}

	err = token.SetTokenToCookie(ctx, token.GenerateTokenRequest{
		Email:      claims.Email,
		JwtKey:     "",
		RefreshKey: "",
	})
	if err != nil {
		return nil, err
	}

	return &desc.RefreshResponse{}, nil
}
