package auth

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	"medblogers_base/internal/pkg/token"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Logout выход из аккаунта
func (i *Implementation) Logout(ctx context.Context, req *desc.LogoutRequest) (*desc.LogoutResponse, error) {
	deleteCookie := &http.Cookie{
		Name:     token.RefreshCookie,
		Value:    "",
		Expires:  time.Now().UTC().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	// Устанавливаем cookie через метаданные
	if err := grpc.SetHeader(ctx, metadata.Pairs(
		"Set-Cookie", deleteCookie.String(),
	)); err != nil {
		return nil, status.Error(codes.Internal, "Failed to set cookie header")
	}

	return &desc.LogoutResponse{}, nil
}
