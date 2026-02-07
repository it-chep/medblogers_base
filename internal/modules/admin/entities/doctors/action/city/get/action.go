package get

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/city/get/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetCities(ctx context.Context) ([]*city.City, error)
}

// Action получение городов
type Action struct {
	actionDal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]*city.City, error) {
	return a.actionDal.GetCities(ctx)
}
