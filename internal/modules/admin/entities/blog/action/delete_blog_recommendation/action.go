package delete_blog_recommendation

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/delete_blog_recommendation/dal"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

type Action struct {
	dal *dal.Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, blogID, recommendationBlogID uuid.UUID) error {
	return a.dal.DeleteRecommendation(ctx, blogID, recommendationBlogID)
}
