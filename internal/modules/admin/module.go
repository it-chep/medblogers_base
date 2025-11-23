package admin

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/admin/action"
	"medblogers_base/internal/pkg/http"
	"medblogers_base/internal/pkg/postgres"
)

// Module модуль отвечающий за работу модуля докторов
type Module struct {
	Actions *action.Aggregator
}

func New(httpConns map[string]http.Executor, config config.AppConfig, pool postgres.PoolWrapper) *Module {

	actions := action.NewAggregator(httpConns, config, pool)

	return &Module{
		Actions: actions,
	}
}
