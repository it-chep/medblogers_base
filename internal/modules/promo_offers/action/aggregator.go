package action

import (
	"medblogers_base/internal/modules/promo_offers/action/brand_detail"
	"medblogers_base/internal/modules/promo_offers/action/filter_brands"
	"medblogers_base/internal/modules/promo_offers/action/filter_offers"
	"medblogers_base/internal/modules/promo_offers/action/filter_settings"
	"medblogers_base/internal/modules/promo_offers/action/offer_detail"
	"medblogers_base/internal/pkg/postgres"
)

type Aggregator struct {
	FilterOffers *filter_offers.Action
	FilterBrands *filter_brands.Action
	GetSettings  *filter_settings.Action
	BrandDetail  *brand_detail.Action
	OfferDetail  *offer_detail.Action
}

func NewAggregator(pool postgres.PoolWrapper) *Aggregator {
	return &Aggregator{
		FilterOffers: filter_offers.New(pool),
		FilterBrands: filter_brands.New(pool),
		GetSettings:  filter_settings.New(pool),
		BrandDetail:  brand_detail.New(pool),
		OfferDetail:  offer_detail.New(pool),
	}
}
