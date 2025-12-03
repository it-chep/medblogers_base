package blogs

import (
	"medblogers_base/internal/modules/blogs/action"
	"medblogers_base/internal/pkg/postgres"
)

// Module модуль отвечающий за работу модуля статей
type Module struct {
	Actions *action.Aggregator
}

func NewModule(pool postgres.PoolWrapper) *Module {
	return &Module{
		Actions: action.NewAggregator(pool),
	}
}
