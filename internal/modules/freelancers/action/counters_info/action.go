package counters_info

import (
	"medblogers_base/internal/modules/doctors/action/counters_info/service/counters"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение настроек главной страницы
type Action struct {
	counters *counters.Service
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		counters: counters.NewService(
			doctor_dal.NewRepository(pool),
		),
	}
}
