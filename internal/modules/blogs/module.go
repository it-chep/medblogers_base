package blogs

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/blogs/action"
	"medblogers_base/internal/modules/blogs/client"
	"medblogers_base/internal/pkg/postgres"
)

// Module модуль отвечающий за работу модуля статей
type Module struct {
	Actions *action.Aggregator
}

func NewModule(pool postgres.PoolWrapper, cfg config.AppConfig) *Module {
	clients := client.NewAggregator(cfg)
	return &Module{
		Actions: action.NewAggregator(pool, clients),
	}
}
