package get_by_id

import (
	"context"

	"github.com/google/uuid"

	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetOfferByID(ctx context.Context, offerID uuid.UUID) (*offerDomain.Offer, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: commonDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, offerID uuid.UUID) (*offerDomain.Offer, error) {
	return a.actionDal.GetOfferByID(ctx, offerID)
}
