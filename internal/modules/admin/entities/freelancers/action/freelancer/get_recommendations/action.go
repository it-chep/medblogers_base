package get_recommendations

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_recommendations/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_recommendations/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_recommendations/service/recommendation"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	recommendation *recommendation.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		recommendation: recommendation.New(dal.NewRepository(pool)),
	}
}

// Do получение рекомендаций для админки
func (a *Action) Do(ctx context.Context, freelancerID int64) ([]dto.Recommendation, error) {
	return a.recommendation.GetRecommendations(ctx, freelancerID)
}
