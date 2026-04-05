package update

import (
	"context"

	"github.com/google/uuid"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/update/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/update/dto"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
)

type CommonDal interface {
	GetOfferByID(ctx context.Context, offerID uuid.UUID) (*offerDomain.Offer, error)
}

type ActionDal interface {
	UpdateOffer(ctx context.Context, offerID uuid.UUID, req dto.UpdateRequest) error
	ReplaceOfferSocialNetworks(ctx context.Context, offerID uuid.UUID, socialNetworkIDs []int64) error
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

func (a *Action) Do(ctx context.Context, offerID uuid.UUID, req dto.UpdateRequest) error {
	if _, err := a.commonDal.GetOfferByID(ctx, offerID); err != nil {
		return err
	}

	return transaction.Exec(ctx, func(ctx context.Context) error {
		// todo audit log

		if err := a.actionDal.UpdateOffer(ctx, offerID, req); err != nil {
			return err
		}

		return a.actionDal.ReplaceOfferSocialNetworks(ctx, offerID, req.SocialNetworkIDs)
	})
}
