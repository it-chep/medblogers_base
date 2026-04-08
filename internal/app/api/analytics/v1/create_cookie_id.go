package v1

import (
	"context"
	"crypto/subtle"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "medblogers_base/internal/pb/medblogers_base/api/analytics/v1"
)

// CreateCookieID создает cookie_id для анонимного пользователя.
func (i *Implementation) CreateCookieID(ctx context.Context, req *desc.CreateCookieIDRequest) (*desc.CreateCookieIDResponse, error) {
	expectedAPIKey := os.Getenv("COOKIE_CREATOR_API_KEY")
	if expectedAPIKey == "" || req.GetToken() == "" ||
		subtle.ConstantTimeCompare([]byte(req.GetToken()), []byte(expectedAPIKey)) != 1 {
		return nil, status.Error(codes.PermissionDenied, "invalid cookie creator api key")
	}

	cookieID, err := i.analytics.Actions.CreateCookieID.Do(ctx, req.GetDomain())
	if err != nil {
		return nil, err
	}

	return &desc.CreateCookieIDResponse{
		CookieId: cookieID.String(),
	}, nil
}
