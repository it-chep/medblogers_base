package get

import (
	"context"
	commondal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetFreelancers(ctx context.Context) ([]*freelancer.Freelancer, error)
}

type Action struct {
	commonDal CommonDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: commondal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]*freelancer.Freelancer, error) {
	return a.commonDal.GetFreelancers(ctx)
}
