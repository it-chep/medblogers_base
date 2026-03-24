package update

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/update/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/update/dto"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
)

type CommonDal interface {
	GetBrandByID(ctx context.Context, brandID int64) (*brandDomain.Brand, error)
}

type ActionDal interface {
	UpdateBrand(ctx context.Context, brandID int64, req dto.UpdateRequest) error
	ReplaceBrandSocialNetworks(ctx context.Context, brandID int64, items []dto.SocialNetworkInput) error
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

func (a *Action) Do(ctx context.Context, brandID int64, req dto.UpdateRequest) error {
	if _, err := a.commonDal.GetBrandByID(ctx, brandID); err != nil {
		return err
	}

	return transaction.Exec(ctx, func(ctx context.Context) error {
		if err := a.actionDal.UpdateBrand(ctx, brandID, req); err != nil {
			return err
		}

		return a.actionDal.ReplaceBrandSocialNetworks(ctx, brandID, req.SocialNetworks)
	})
}
