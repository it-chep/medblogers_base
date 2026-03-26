package brand

import (
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/activate"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/add_network"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/create"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/deactivate"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/delete_network"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get_by_id"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/save_brand_photo"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/update"
	"medblogers_base/internal/pkg/postgres"
)

type PromoOfferBrandAggregator struct {
	CreateBrand     *create.Action
	GetBrands       *get.Action
	GetBrandByID    *get_by_id.Action
	UpdateBrand     *update.Action
	AddNetwork      *add_network.Action
	DeleteNetwork   *delete_network.Action
	SaveBrandPhoto  *save_brand_photo.Action
	ActivateBrand   *activate.Action
	DeactivateBrand *deactivate.Action
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *PromoOfferBrandAggregator {
	return &PromoOfferBrandAggregator{
		CreateBrand:     create.New(pool),
		GetBrands:       get.New(clients, pool),
		GetBrandByID:    get_by_id.New(clients, pool),
		UpdateBrand:     update.New(pool),
		AddNetwork:      add_network.New(pool),
		DeleteNetwork:   delete_network.New(pool),
		SaveBrandPhoto:  save_brand_photo.New(clients, pool),
		ActivateBrand:   activate.New(pool),
		DeactivateBrand: deactivate.New(pool),
	}
}
