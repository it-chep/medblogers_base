package search

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/search/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	SearchFreelancers(ctx context.Context, query string) ([]*freelancer.Freelancer, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, query string) ([]*freelancer.Freelancer, error) {
	return a.actionDal.SearchFreelancers(ctx, query)
}
