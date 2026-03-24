package get_cooperation_types

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/get_cooperation_types/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetCooperationTypes(ctx context.Context) (dictionary.NamedItems, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: actionDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) (dictionary.NamedItems, error) {
	return a.actionDal.GetCooperationTypes(ctx)
}
