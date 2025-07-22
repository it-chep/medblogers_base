package doctors

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/doctors/action"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/postgres"
)

// Module модуль отвечающий за работу модуля докторов
type Module struct {
	Actions *action.Aggregator
}

func New(config *config.Config, pool postgres.PoolWrapper) *Module {
	clients := client.NewAggregator(config)

	actions := action.NewAggregator(clients, pool, config)

	return &Module{
		Actions: actions,
	}
}
