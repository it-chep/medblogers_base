package create

import (
	"context"

	"github.com/google/uuid"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/create/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/create/dto"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
)

type ActionDal interface {
	CreateOffer(ctx context.Context, req dto.CreateRequest) (uuid.UUID, error)
	ReplaceOfferSocialNetworks(ctx context.Context, offerID uuid.UUID, socialNetworkIDs []int64) error
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{actionDal: actionDal.NewRepository(pool)}
}

func (a *Action) Do(ctx context.Context, req dto.CreateRequest) (offerID uuid.UUID, err error) {
	err = transaction.Exec(ctx, func(ctx context.Context) error {
		offerID, err = a.actionDal.CreateOffer(ctx, req)
		if err != nil {
			return err
		}

		return a.actionDal.ReplaceOfferSocialNetworks(ctx, offerID, req.SocialNetworkIDs)
	})

	return offerID, err
}
