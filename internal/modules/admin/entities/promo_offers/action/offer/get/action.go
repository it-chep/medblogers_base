package get

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get/dal"
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetOffers(ctx context.Context) (offerDomain.Offers, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: actionDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) (offerDomain.Offers, error) {
	return a.actionDal.GetOffers(ctx)
}
