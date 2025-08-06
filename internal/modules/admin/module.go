package admin

import "medblogers_base/internal/modules/admin/action"

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
