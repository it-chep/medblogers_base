package update_draft_blog

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/action/blog/action/update_draft_blog/dal"
	"medblogers_base/internal/modules/admin/action/blog/action/update_draft_blog/dto"
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
func (a *Action) Do(ctx context.Context, blogID uuid.UUID, req dto.Request) error {
	blog, err := a.dal.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	if blog.IsActive.Bool {
		return errors.New("Статью надо сначала снять с публикации чтобы поменять")
	}

	return a.dal.UpdateBlog(ctx, blogID, req)
}
