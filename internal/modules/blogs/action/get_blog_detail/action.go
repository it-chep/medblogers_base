package get_blog_detail

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail/dal"
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение карточки статьи
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
func (a *Action) Do(ctx context.Context, slug string) (*blog.Blog, error) {
	return a.dal.GetBlogDetail(ctx, slug)
}
