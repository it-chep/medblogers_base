package promo_offers

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/promo_offers/action"
	"medblogers_base/internal/modules/promo_offers/client"
	"medblogers_base/internal/pkg/postgres"
)

type Module struct {
	Actions *action.Aggregator
}

func New(pool postgres.PoolWrapper, cfg config.AppConfig) *Module {
	clients := client.NewAggregator(cfg)

	return &Module{
		Actions: action.NewAggregator(pool, clients),
	}
}
