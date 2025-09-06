package action

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	PreliminaryFilterCount *preliminary_filter_count.Action
}

// NewAggregator конструктор
func NewAggregator(clients *client.Aggregator, pool postgres.PoolWrapper, config config.AppConfig) *Aggregator {
	return &Aggregator{
		PreliminaryFilterCount: preliminary_filter_count.New(pool),
	}
}
