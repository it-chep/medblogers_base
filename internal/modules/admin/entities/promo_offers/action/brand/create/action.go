package create

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/create/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/create/dto"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
)

type ActionDal interface {
	CreateBrand(ctx context.Context, req dto.CreateRequest) (int64, error)
	ReplaceBrandSocialNetworks(ctx context.Context, brandID int64, items []dto.SocialNetworkInput) error
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{actionDal: actionDal.NewRepository(pool)}
}

func (a *Action) Do(ctx context.Context, req dto.CreateRequest) (brandID int64, err error) {
	err = transaction.Exec(ctx, func(ctx context.Context) error {
		brandID, err = a.actionDal.CreateBrand(ctx, req)
		if err != nil {
			return err
		}

		return a.actionDal.ReplaceBrandSocialNetworks(ctx, brandID, req.SocialNetworks)
	})

	return brandID, err
}
