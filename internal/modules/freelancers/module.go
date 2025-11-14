package freelancers

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/freelancers/action"
	"medblogers_base/internal/modules/freelancers/client"
	"medblogers_base/internal/pkg/http"
	"medblogers_base/internal/pkg/postgres"
)

// Module модуль отвечающий за работу модуля докторов
type Module struct {
	Actions *action.Aggregator
}

func New(httpConns map[string]http.Executor, config config.AppConfig, pool postgres.PoolWrapper) *Module {
	clients := client.NewAggregator(httpConns, config)

	actions := action.NewAggregator(clients, pool, config)

	return &Module{
		Actions: actions,
	}
}
