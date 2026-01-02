package delete

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/city/delete/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	DeleteCity(ctx context.Context, id int64) error
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

func (a *Action) Do(ctx context.Context, id int64) error {
	return a.actionDal.DeleteCity(ctx, id)
}
