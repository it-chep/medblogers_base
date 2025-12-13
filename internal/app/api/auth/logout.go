package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	desc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	"medblogers_base/internal/pkg/token"
)

// Logout выход из аккаунта
func (i *Implementation) Logout(ctx context.Context, req *desc.LogoutRequest) (*desc.LogoutResponse, error) {
	md := metadata.New(
		map[string]string{
			"set-cookie": fmt.Sprintf(
				"%s=%s; Path=/; Max-Age=%d; HttpOnly; SameSite=Lax",
				token.RefreshCookie, "", -1,
			),
		},
	)

	// Send metadata as trailer (Clay should convert this to HTTP headers)
	if err := grpc.SetHeader(ctx, md); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}

	return &desc.LogoutResponse{}, nil
}
