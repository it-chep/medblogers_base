package get_categories

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/get_categories/dal"
	"medblogers_base/internal/modules/admin/action/blog/action/get_categories/dto"
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

func (a *Action) Do(ctx context.Context) (dto.Categories, error) {
	return a.dal.GetCategories(ctx)
}
