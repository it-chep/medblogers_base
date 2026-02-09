package search

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/social_network/search/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/social_network"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	SearchNetworks(ctx context.Context, query string) (social_network.Networks, error)
}
type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, query string) (social_network.Networks, error) {
	return a.actionDal.SearchNetworks(ctx, query)
}
