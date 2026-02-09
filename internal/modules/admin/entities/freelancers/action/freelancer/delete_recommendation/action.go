package delete_recommendation

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete_recommendation/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	DeleteRecommendation(ctx context.Context, freelancerID, doctorID int64) error
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID, doctorID int64) error {
	return a.actionDal.DeleteRecommendation(ctx, freelancerID, doctorID)
}
