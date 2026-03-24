package brand

import (
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/activate"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/create"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/deactivate"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get_by_id"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/update"
	"medblogers_base/internal/pkg/postgres"
)

type PromoOfferBrandAggregator struct {
	CreateBrand     *create.Action
	GetBrands       *get.Action
	GetBrandByID    *get_by_id.Action
	UpdateBrand     *update.Action
	ActivateBrand   *activate.Action
	DeactivateBrand *deactivate.Action
}

func New(pool postgres.PoolWrapper) *PromoOfferBrandAggregator {
	return &PromoOfferBrandAggregator{
		CreateBrand:     create.New(pool),
		GetBrands:       get.New(pool),
		GetBrandByID:    get_by_id.New(pool),
		UpdateBrand:     update.New(pool),
		ActivateBrand:   activate.New(pool),
		DeactivateBrand: deactivate.New(pool),
	}
}
