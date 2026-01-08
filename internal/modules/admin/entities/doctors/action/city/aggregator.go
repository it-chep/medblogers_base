package city

import (
	"medblogers_base/internal/modules/admin/entities/doctors/action/city/create"
	delete_city "medblogers_base/internal/modules/admin/entities/doctors/action/city/delete"
	"medblogers_base/internal/modules/admin/entities/doctors/action/city/get"
	"medblogers_base/internal/modules/admin/entities/doctors/action/city/search"
	"medblogers_base/internal/pkg/postgres"
)

type DoctorCityAggregator struct {
	CreateCity   *create.Action
	DeleteCity   *delete_city.Action
	GetCities    *get.Action
	SearchCities *search.Action
}

func New(pool postgres.PoolWrapper) *DoctorCityAggregator {
	return &DoctorCityAggregator{
		CreateCity:   create.New(pool),
		DeleteCity:   delete_city.New(pool),
		GetCities:    get.New(pool),
		SearchCities: search.New(pool),
	}
}
