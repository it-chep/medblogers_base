package deactivate

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/deactivate/dal"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetOfferByID(ctx context.Context, offerID uuid.UUID) (*offerDomain.Offer, error)
}

type ActionDal interface {
	DeactivateOffer(ctx context.Context, offerID uuid.UUID) error
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

func (a *Action) Do(ctx context.Context, offerID uuid.UUID) error {
	item, err := a.commonDal.GetOfferByID(ctx, offerID)
	if err != nil {
		return err
	}

	if !item.GetIsActive() {
		return errors.New("Оффер уже неактивен")
	}

	return a.actionDal.DeactivateOffer(ctx, offerID)
}
