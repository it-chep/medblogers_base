package get_doctor_additional_cities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_doctor_additional_cities/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_doctor_additional_cities/service/cities"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	cities *cities.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		cities: cities.New(dal.NewRepository(pool)),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64) ([]*city.City, error) {
	return a.cities.GetAdditionalCities(ctx, doctorID)
}
