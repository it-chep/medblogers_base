package doctors

import (
	"github.com/it-chep/medblogers_base/internal/modules/doctors/action"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/client"
	"github.com/it-chep/medblogers_base/internal/pkg/postgres"
)

// Module модуль отвечающий за работу модуля докторов
type Module struct {
	Actions *action.Aggregator
}

func New(pool postgres.PoolWrapper) *Module {
	clients := client.NewAggregator()

	actions := action.NewAggregator(clients, pool)

	return &Module{
		Actions: actions,
	}
}
