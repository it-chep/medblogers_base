package get_cooperation_types

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_cooperation_types/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type Dal interface {
	GetFreelancerCooperationTypes(ctx context.Context) ([]*freelancer.CooperationType, error)
}

type Action struct {
	dal Dal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]*freelancer.CooperationType, error) {
	return a.dal.GetFreelancerCooperationTypes(ctx)
}
