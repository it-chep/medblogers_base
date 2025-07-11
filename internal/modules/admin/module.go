package admin

import "github.com/it-chep/medblogers_base/internal/modules/admin/action"

// Module модуль отвечающий за работу админки
type Module struct {
	Actions *action.Aggregator
}

func New() *Module {
	actions := action.NewAggregator()

	return &Module{
		Actions: actions,
	}
}
