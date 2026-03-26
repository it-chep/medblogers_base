package offer

import (
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/a
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/activate"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/create"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/deactivate"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/filter"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get_by_id"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/postgres"
)

type PromoOfferOfferAggregator struct {
	CreateOffer     *create.Action
	GetOffers       *get.Action
	FilterOffers    *filter.Action
	GetOfferByID    *get_by_id.Action
	UpdateOffer     *update.Action
	ActivateOffer   *activate.Action
	DeactivateOffer *deactivate.Action
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *PromoOfferOfferAggregator {
	return &PromoOfferOfferAggregator{
		CreateOffer:     create.New(pool),
		GetOffers:       get.New(clients, pool),
		FilterOffers:    filter.New(clients, pool),
		GetOfferByID:    get_by_id.New(clients, pool),
		UpdateOffer:     update.New(pool),
		ActivateOffer:   activate.New(pool),
		DeactivateOffer: deactivate.New(pool),
	}
}
