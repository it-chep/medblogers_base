package auth

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	"medblogers_base/internal/pkg/token"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Email and password are required")
	}

	user, err := i.auth.Actions.GetUserInfo.Do(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.GetPassword())); err != nil {
		return nil, status.Error(codes.Internal, "Please try again later")
	}

	err = token.SetTokenToCookie(ctx, token.GenerateTokenRequest{Email: user.GetEmail(), JwtKey: i.config.JWTConfig.Secret, RefreshKey: i.config.JWTConfig.RefreshSecret})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}

	return &desc.LoginResponse{}, nil
}
