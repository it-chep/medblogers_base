package action

import (
	"medblogers_base/internal/config"
	blog_action "medblogers_base/internal/modules/admin/action/blog/action"
	mm_action "medblogers_base/internal/modules/admin/action/mm/action"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/http"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	BlogModule *blog_action.BlogModuleAggregator
	MMModule   *mm_action.MMActionAggregator
}

func NewAggregator(httpConns map[string]http.Executor, config config.AppConfig, pool postgres.PoolWrapper) *Aggregator {

	clients := client.NewAggregator(httpConns, config)

	return &Aggregator{
		MMModule:   mm_action.New(pool, clients, config),
		BlogModule: blog_action.New(pool, clients),
	}
}
