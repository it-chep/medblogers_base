package get_doctor_additional_specialities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_doctor_additional_specialities/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_doctor_additional_specialities/service/specialities"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	specialities *specialities.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		specialities: specialities.New(dal.NewRepository(pool)),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64) ([]*speciality.Speciality, error) {
	return a.specialities.GetAdditionalSpecialities(ctx, doctorID)
}
