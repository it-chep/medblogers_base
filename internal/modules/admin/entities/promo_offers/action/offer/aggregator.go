package offer

import (
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/activate"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/create"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/deactivate"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get_by_id"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/update"
	"medblogers_base/internal/pkg/postgres"
)

type PromoOfferOfferAggregator struct {
	CreateOffer     *create.Action
	GetOffers       *get.Action
	GetOfferByID    *get_by_id.Action
	UpdateOffer     *update.Action
	ActivateOffer   *activate.Action
	DeactivateOffer *deactivate.Action
}

func New(pool postgres.PoolWrapper) *PromoOfferOfferAggregator {
	return &PromoOfferOfferAggregator{
		CreateOffer:     create.New(pool),
		GetOffers:       get.New(pool),
		GetOfferByID:    get_by_id.New(pool),
		UpdateOffer:     update.New(pool),
		ActivateOffer:   activate.New(pool),
		DeactivateOffer: deactivate.New(pool),
	}
}
