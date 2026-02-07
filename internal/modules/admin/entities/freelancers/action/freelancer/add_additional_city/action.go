package add_additional_city

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_additional_city/dal"
	commondal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
}

type ActionDal interface {
	AddAdditionalCity(ctx context.Context, freelancerID, cityID int64) error
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
		commonDal: commondal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID, cityID int64) error {
	_, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return err
	}

	return a.actionDal.AddAdditionalCity(ctx, freelancerID, cityID)
}
