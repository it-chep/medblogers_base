package settings

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/settings/action"
	"medblogers_base/internal/pkg/postgres"
)

type Module struct {
	Actions *action.Aggregator
}

func New(pool postgres.PoolWrapper, cfg config.AppConfig) *Module {
	return &Module{
		Actions: action.New(pool, cfg),
	}
}
