package unpublish_blog

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/unpublish_blog/dal"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
	"github.com/pkg/errors"
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

// Do снятие статьи с публикации
func (a *Action) Do(ctx context.Context, blogID uuid.UUID) error {
	blog, err := a.dal.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	if !blog.IsActive.Bool {
		return errors.New("Статья уже снята с публикации")
	}

	return a.dal.UnPublishBlog(ctx, blogID)
}
