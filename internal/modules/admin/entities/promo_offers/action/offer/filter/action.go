package filter

import (
	"context"

	"medblogers_base/internal/modules/admin/client"
	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/filter/dal"
	filterDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/filter/dto"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get"
	getDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get/d
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	FilterOffers(ctx context.Context, req filterDTO.Request) (offerDomain.Offers, error)
}

type Action struct {
	actionDal ActionDal
	getAction *get.Action
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: actionDal.NewRepository(pool),
		getAction: get.New(clients, pool),
	}
}

func (a *Action) Do(ctx context.Context, req filterDTO.Request) ([]getDTO.Offer, error) {
	offers, err := a.actionDal.FilterOffers(ctx, req)
	if err != nil {
		return nil, err
	}

	return a.getAction.EnrichOffers(ctx, offers)
}
