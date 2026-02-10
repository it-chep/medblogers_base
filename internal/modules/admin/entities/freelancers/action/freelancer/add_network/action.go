package add_network

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_network/dal"
	commondal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
}

type ActionDal interface {
	AddNetwork(ctx context.Context, freelancerID, networkID int64) error
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
		commonDal: commondal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID, networkID int64) error {
	_, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return err
	}

	return a.actionDal.AddNetwork(ctx, freelancerID, networkID)
}
