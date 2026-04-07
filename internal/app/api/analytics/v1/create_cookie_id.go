package v1

import (
	"context"

	desc "medblogers_base/internal/pb/medblogers_base/api/analytics/v1"
)

// CreateCookieID создает cookie_id для анонимного пользователя.
func (i *Implementation) CreateCookieID(ctx context.Context, req *desc.CreateCookieIDRequest) (*desc.CreateCookieIDResponse, error) {
	cookieID, err := i.analytics.Actions.CreateCookieID.Do(ctx, req.GetDomain())
	if err != nil {
		return nil, err
	}

	return &desc.CreateCookieIDResponse{
		CookieId: cookieID.String(),
	}, nil
}
