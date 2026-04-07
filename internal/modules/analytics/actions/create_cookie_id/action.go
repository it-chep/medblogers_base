package create_cookie_id

import (
	"context"
	"medblogers_base/internal/modules/analytics/actions/create_cookie_id/dal"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

// Action создание cookie_id.
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
func (a *Action) Do(ctx context.Context, domain string) (uuid.UUID, error) {
	cookieID := uuid.New()

	if err := a.dal.CreateCookieUser(ctx, cookieID, domain); err != nil {
		return uuid.Nil, err
	}

	return cookieID, nil
}
