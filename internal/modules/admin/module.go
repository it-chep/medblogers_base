package admin

import (
	"medblogers_base/internal/modules/admin/action"
	"medblogers_base/internal/pkg/postgres"
)

// Module модуль отвечающий за работу модуля докторов
type Module struct {
	Actions *action.Aggregator
}

func New(pool postgres.PoolWrapper) *Module {

	actions := action.NewAggregator(pool)

	return &Module{
		Actions: actions,
	}
}
