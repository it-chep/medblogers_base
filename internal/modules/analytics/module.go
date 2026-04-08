package analytics

import (
	"medblogers_base/internal/modules/analytics/actions"
	"medblogers_base/internal/pkg/postgres"
)

// Module модуль аналитики.
type Module struct {
	Actions *actions.Aggregator
}

// New .
func New(pool postgres.PoolWrapper) *Module {
	return &Module{
		Actions: actions.NewAggregator(pool),
	}
}
