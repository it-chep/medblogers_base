package auth

import (
	"medblogers_base/internal/modules/auth/action"
	"medblogers_base/internal/pkg/postgres"
)

// Module модуль отвечающий за работу модуля утентификации
type Module struct {
	Actions *action.Aggregator
}

func New(pool postgres.PoolWrapper) *Module {

	actions := action.NewAggregator(pool)

	return &Module{
		Actions: actions,
	}
}
