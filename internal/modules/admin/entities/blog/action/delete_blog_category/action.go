package delete_blog_category

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/delete_blog_category/dal"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
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

func (a *Action) Do(ctx context.Context, blogID uuid.UUID, categoryID int64) error {
	return a.dal.DeleteCategory(ctx, blogID, categoryID)
}
