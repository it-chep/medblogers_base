package action

import (
	"medblogers_base/internal/modules/doctors/action/create_doctor"
	"medblogers_base/internal/modules/doctors/action/doctor_detail"
	"medblogers_base/internal/modules/doctors/action/doctors_filter"
	"medblogers_base/internal/modules/doctors/action/doctors_list"
	"medblogers_base/internal/modules/doctors/action/search_doctor"
	"medblogers_base/internal/modules/doctors/action/settings"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	CreateDoctor  *create_doctor.Action
	DoctorDetail  *doctor_detail.Action
	DoctorsFilter *doctors_filter.Action
	DoctorsList   *doctors_list.Action
	SearchDoctor  *search_doctor.Action
	Settings      *settings.Action
}

// NewAggregator конструктор
func NewAggregator() *Aggregator {
	return &Aggregator{
		CreateDoctor:  create_doctor.New(),
		DoctorDetail:  doctor_detail.New(),
		DoctorsFilter: doctors_filter.New(),
		DoctorsList:   doctors_list.New(),
		SearchDoctor:  search_doctor.New(),
		Settings:      settings.New(),
	}
}
