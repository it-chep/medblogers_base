package action

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/doctors/action/counters_info"
	"medblogers_base/internal/modules/doctors/action/create_doctor"
	"medblogers_base/internal/modules/doctors/action/doctor_detail"
	"medblogers_base/internal/modules/doctors/action/doctors_filter"
	"medblogers_base/internal/modules/doctors/action/get_all_cities"
	"medblogers_base/internal/modules/doctors/action/get_all_specialities"
	"medblogers_base/internal/modules/doctors/action/search_doctor"
	"medblogers_base/internal/modules/doctors/action/settings"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	CreateDoctor    *create_doctor.Action
	CounterInfo     *counters_info.Action
	DoctorDetail    *doctor_detail.Action
	DoctorsFilter   *doctors_filter.Action
	SearchDoctor    *search_doctor.Action
	Settings        *settings.Action
	AllCities       *get_all_cities.Action
	AllSpecialities *get_all_specialities.Action
}

// NewAggregator конструктор
func NewAggregator(clients *client.Aggregator, pool postgres.PoolWrapper, config *config.Config) *Aggregator {
	return &Aggregator{
		CreateDoctor:    create_doctor.New(clients, pool, config),
		CounterInfo:     counters_info.New(clients, pool),
		DoctorDetail:    doctor_detail.New(clients, pool),
		DoctorsFilter:   doctors_filter.New(clients, pool),
		SearchDoctor:    search_doctor.New(clients, pool),
		Settings:        settings.New(clients, pool),
		AllCities:       get_all_cities.New(pool),
		AllSpecialities: get_all_specialities.New(pool),
	}
}
