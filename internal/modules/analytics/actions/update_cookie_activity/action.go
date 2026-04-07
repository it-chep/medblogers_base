package update_cookie_activity

import (
	"context"
	"medblogers_base/internal/modules/analytics/actions/update_cookie_activity/dal"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

// Action обновление активности cookie пользователя.
type Action struct {
	dal *dal.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do .
func (a *Action) Do(ctx context.Context, cookieID string) error {
	if cookieID == "" {
		return nil
	}

	parsedCookieID, err := uuid.Parse(cookieID)
	if err != nil {
		return nil
	}

	return a.dal.UpdateLastActivity(ctx, parsedCookieID)
}
