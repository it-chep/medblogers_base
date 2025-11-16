package auth

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	"medblogers_base/internal/pkg/token"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {
	user, err := i.auth.Actions.Register.Do(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	err = token.SetTokenToCookie(ctx, token.GenerateTokenRequest{Email: user.GetEmail(), JwtKey: "i.jwt.JwtSecret", RefreshKey: "i.jwt.RefreshSecret"}) // todo надо доставать это из конфига
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}

	return &desc.RegisterResponse{}, nil
}
