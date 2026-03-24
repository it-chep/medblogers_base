package get_by_id

import (
	"context"

	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetBrandByID(ctx context.Context, brandID int64) (*brandDomain.Brand, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: commonDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, brandID int64) (*brandDomain.Brand, error) {
	return a.actionDal.GetBrandByID(ctx, brandID)
}
