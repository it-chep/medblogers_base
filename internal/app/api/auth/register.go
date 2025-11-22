package auth

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	"medblogers_base/internal/pkg/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {
	err := i.auth.Actions.Register.Do(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	err = token.SetTokenToCookie(ctx, token.GenerateTokenRequest{Email: req.GetEmail(), JwtKey: i.config.JWTConfig.Secret, RefreshKey: i.config.JWTConfig.RefreshSecret})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}

	return &desc.RegisterResponse{}, nil
}
