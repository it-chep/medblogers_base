package city

import (
	"medblogers_base/internal/modules/admin/entities/doctors/action/speciality/create"
	delete_speciality "medblogers_base/internal/modules/admin/entities/doctors/action/speciality/delete"
	"medblogers_base/internal/modules/admin/entities/doctors/action/speciality/get"
	"medblogers_base/internal/modules/admin/entities/doctors/action/speciality/search"
	"medblogers_base/internal/pkg/postgres"
)

type DoctorSpecialityAggregator struct {
	CreateSpeciality   *create.Action
	DeleteSpeciality   *delete_speciality.Action
	GetSpecialities    *get.Action
	SearchSpecialities *search.Action
}

func New(pool postgres.PoolWrapper) *DoctorSpecialityAggregator {
	return &DoctorSpecialityAggregator{
		CreateSpeciality:   create.New(pool),
		DeleteSpeciality:   delete_speciality.New(pool),
		GetSpecialities:    get.New(pool),
		SearchSpecialities: search.New(pool),
	}
}
