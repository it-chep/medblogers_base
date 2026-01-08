package speciality

import (
	"medblogers_base/internal/modules/admin/entities/freelancers/action/speciality/create"
	delete_speciality "medblogers_base/internal/modules/admin/entities/freelancers/action/speciality/delete"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/speciality/get"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/speciality/search"
	"medblogers_base/internal/pkg/postgres"
)

type FreelancerSpecialityAggregator struct {
	CreateSpeciality   *create.Action
	DeleteSpeciality   *delete_speciality.Action
	GetSpecialities    *get.Action
	SearchSpecialities *search.Action
}

func New(pool postgres.PoolWrapper) *FreelancerSpecialityAggregator {
	return &FreelancerSpecialityAggregator{
		CreateSpeciality:   create.New(pool),
		DeleteSpeciality:   delete_speciality.New(pool),
		GetSpecialities:    get.New(pool),
		SearchSpecialities: search.New(pool),
	}
}
