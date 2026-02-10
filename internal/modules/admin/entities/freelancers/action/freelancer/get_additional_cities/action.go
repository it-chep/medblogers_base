package get_additional_cities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_additional_cities/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
	"medblogers_base/internal/pkg/postgres"
)

type Dal interface {
	GetAdditionalCities(ctx context.Context, freelancerID int64) ([]*city.City, error)
}

type Action struct {
	dal Dal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID int64) ([]*city.City, error) {
	return a.dal.GetAdditionalCities(ctx, freelancerID)
}
