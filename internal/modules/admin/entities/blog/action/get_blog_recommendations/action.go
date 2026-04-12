package get_blog_recommendations

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blog_recommendations/dal"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blog_recommendations/dto"
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

func (a *Action) Do(ctx context.Context, blogID uuid.UUID) ([]dto.Recommendation, error) {
	return a.dal.GetRecommendations(ctx, blogID)
}
