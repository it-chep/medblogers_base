package action

import (
	"medblogers_base/internal/modules/admin/action/create_doctor"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	CreateDoctor *create_doctor.Action
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		CreateDoctor: create_doctor.New(),
	}
}
