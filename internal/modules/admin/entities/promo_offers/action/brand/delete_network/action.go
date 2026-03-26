package delete_network

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/delete_network/dal"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetBrandByID(ctx context.Context, brandID int64) (*brandDomain.Brand, error)
}

type ActionDal interface {
	DeleteNetwork(ctx context.Context, brandID, socialNetworkID int64) error
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: commonDal.NewRepository(pool),
		actionDal: actionDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, brandID, socialNetworkID int64) error {
	if _, err := a.commonDal.GetBrandByID(ctx, brandID); err != nil {
		return err
	}

	return a.actionDal.DeleteNetwork(ctx, brandID, socialNetworkID)
}
