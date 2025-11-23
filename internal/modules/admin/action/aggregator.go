package action

import "medblogers_base/internal/pkg/postgres"

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	//CreateDoctor *create_doctor.Action
}

func NewAggregator(pool postgres.PoolWrapper) *Aggregator {
	return &Aggregator{
		//CreateDoctor: create_doctor.New(),
	}
}
