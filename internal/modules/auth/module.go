package auth

import "medblogers_base/internal/modules/auth/action"

// Module модуль отвечающий за работу модуля докторов
type Module struct {
	Actions *action.Aggregator
}

func NewModule() *Module {
	return &Module{}
}
