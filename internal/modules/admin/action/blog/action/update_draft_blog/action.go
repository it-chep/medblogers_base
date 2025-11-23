package update_draft_blog

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/update_draft_blog/dal"
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

// Do обновление статьи
func (a *Action) Do(ctx context.Context) error {
	return nil
}
