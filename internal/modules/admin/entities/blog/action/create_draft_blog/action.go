package create_draft_blog

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/create_draft_blog/dal"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"

	"github.com/google/uuid"
)

type ActionDal interface {
	CreateDraftBlogs(ctx context.Context, title, slug string, id uuid.UUID) error
	CreateBlogBreadcrumb(ctx context.Context, title, slug string) error
}

// Action создание драфтовой статьи
type Action struct {
	dal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do .
func (a *Action) Do(ctx context.Context, title, slug string) (uuid.UUID, error) {
	blogID, _ := uuid.NewV7()

	return blogID, transaction.Exec(ctx, func(ctx context.Context) error {
		if err := a.dal.CreateDraftBlogs(ctx, title, slug, blogID); err != nil {
			return err
		}

		return a.dal.CreateBlogBreadcrumb(ctx, title, slug)
	})
}
