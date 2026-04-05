package create_cooperation_type

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/create_cooperation_type/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	CreateCooperationType(ctx context.Context, name string) (int64, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{actionDal: actionDal.NewRepository(pool)}
}

func (a *Action) Do(ctx context.Context, name string) (int64, error) {
	return a.actionDal.CreateCooperationType(ctx, name)
}
