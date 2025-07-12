package doctors

import (
	"github.com/it-chep/medblogers_base/internal/modules/doctors/action"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/client"
)

// Module модуль отвечающий за работу модуля докторов
type Module struct {
	Actions *action.Aggregator
}

func New() *Module {
	clients := client.NewAggregator()

	actions := action.NewAggregator(clients)

	return &Module{
		Actions: actions,
	}
}
