package get_blogs_categories

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_blogs_categories/dal"
	"medblogers_base/internal/modules/blogs/domain/category"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение списка категорий статей
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
func (a *Action) Do(ctx context.Context) (category.Categories, error) {
	return a.dal.GetAllCategories(ctx)
}
