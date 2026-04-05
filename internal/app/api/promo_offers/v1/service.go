package v1

import (
	"medblogers_base/internal/modules/promo_offers"
	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

type Implementation struct {
	desc.UnimplementedPromoOffersServiceServer

	promoOffers *promo_offers.Module
}

func NewService(module *promo_offers.Module) *Implementation {
	return &Implementation{
		promoOffers: module,
	}
}
