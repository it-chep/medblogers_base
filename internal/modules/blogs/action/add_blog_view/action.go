package add_blog_view

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/add_blog_view/dal"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

// Action добавление просмотра статьи.
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
func (a *Action) Do(ctx context.Context, slug string, cookieID string) error {
	blogEntity, err := a.dal.GetBlogBySlug(ctx, slug)
	if err != nil {
		return err
	}

	isViewed, err := a.dal.IsBlogViewedByCookieForLast7Days(ctx, blogEntity.GetID(), cookieID)
	if err != nil {
		return err
	}

	if isViewed {
		return nil
	}

	viewID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	return a.dal.CreateBlogView(ctx, viewID, blogEntity.GetID(), cookieID)
}
