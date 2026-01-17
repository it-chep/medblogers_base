package seo

import (
	"medblogers_base/internal/modules/seo/action"
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
