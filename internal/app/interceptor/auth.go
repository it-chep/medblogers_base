package interceptor

import (
	"context"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/auth/domain/user"
	pkgctx "medblogers_base/internal/pkg/context"
	"medblogers_base/internal/pkg/token"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CheckPermissions мидлвара, которая проверяет есть ли пермишены на ручку
type CheckPermissions interface {
	Do(ctx context.Context, email, path string) (*user.User, error)
}

// ExecuteWithPermissions - декоратор с проверкой прав
func ExecuteWithPermissions(checker CheckPermissions) func(
	ctx context.Context,
	path string,
	fn func(ctx context.Context) error,
) error {
	return func(ctx context.Context, path string, fn func(ctx context.Context) error) error {
		authNeeded, _ := strconv.ParseBool(os.Getenv("IS_AUTH_NEED"))
		if !authNeeded {
			return fn(ctx)
		}

		email := pkgctx.GetEmailFromContext(ctx)

		if len(email) == 0 {
			return status.Error(codes.PermissionDenied, "permission denied")
		}

		usr, err := checker.Do(ctx, email, path)
		if err != nil {
			return status.Error(codes.PermissionDenied, "permission denied")
		}

		userCtx := context.WithValue(ctx, "user", usr)
		userIDCtx := pkgctx.WithUserIDContext(userCtx, usr.GetID())

		return fn(userIDCtx)
	}
}

func EmailInterceptor(cfg *config.Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}

		authHeaders := md.Get("grpcgateway-authorization")
		if len(authHeaders) == 0 {
			return handler(ctx, req)
		}

		authHeader := authHeaders[0]
		claims, _ := token.AccessClaimsFromRequest(authHeader, cfg.JWTConfig.Secret)
		if claims == nil || claims.Email == "" {
			return handler(ctx, req)
		}

		newCtx := pkgctx.WithEmailContext(ctx, claims.Email)
		return handler(newCtx, req)
	}
}
