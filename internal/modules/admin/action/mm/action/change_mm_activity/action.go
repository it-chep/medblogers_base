package change_mm_activity

import (
	"context"
	"medblogers_base/internal/modules/admin/action/mm/action/change_mm_activity/dal"
	"medblogers_base/internal/pkg/postgres"
)

// Action .
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
func (a *Action) Do(ctx context.Context, mmID int64, activity bool) error {
	return a.dal.ChangeMMActivity(ctx, mmID, activity)
}
