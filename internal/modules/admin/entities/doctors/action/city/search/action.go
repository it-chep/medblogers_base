package search

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/city/search/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	SearchCities(ctx context.Context, name string) ([]*city.City, error)
}

// Action поиск городов
type Action struct {
	actionDal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, name string) ([]*city.City, error) {
	return a.actionDal.SearchCities(ctx, name)
}
