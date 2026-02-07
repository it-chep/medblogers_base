package create

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/city/create/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	CreateCity(ctx context.Context, name string) error
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

func (a *Action) Do(ctx context.Context, name string) error {
	return a.actionDal.CreateCity(ctx, name)
}
