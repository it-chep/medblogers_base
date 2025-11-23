package create_draft_blog

import (
	"context"
	"github.com/google/uuid"
	"medblogers_base/internal/modules/admin/action/blog/action/create_draft_blog/dal"
	"medblogers_base/internal/pkg/postgres"
)

// Action создание драфтовой статьи
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
func (a *Action) Do(ctx context.Context, title string) (uuid.UUID, error) {
	blogID, _ := uuid.NewV7()

	return blogID, a.dal.CreateDraftBlogs(ctx, title, blogID)
}
