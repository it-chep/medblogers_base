package get

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get/dal"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetBrands(ctx context.Context) (brandDomain.Brands, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: actionDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) (brandDomain.Brands, error) {
	return a.actionDal.GetBrands(ctx)
}
