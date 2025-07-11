package doctors

import "github.com/it-chep/medblogers_base/internal/modules/doctors/action"

// Module модуль отвечающий за работу модуля докторов
type Module struct {
	Actions *action.Aggregator
}

func New() *Module {
	actions := action.NewAggregator()

	return &Module{
		Actions: actions,
	}
}
