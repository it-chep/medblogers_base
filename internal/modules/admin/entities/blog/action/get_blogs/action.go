package get_blogs

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blogs/dal"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blogs/dto"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение всех статей
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
func (a *Action) Do(ctx context.Context) (dto.Blogs, error) {
	return a.dal.GetBlogs(ctx)
}
