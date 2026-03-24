package promo_offers

import (
	"medblogers_base/internal/modules/promo_offers/action"
	"medblogers_base/internal/pkg/postgres"
)

type Module struct {
	Actions *action.Aggregator
}

func New(pool postgres.PoolWrapper) *Module {
	return &Module{
		Actions: action.NewAggregator(pool),
	}
}
