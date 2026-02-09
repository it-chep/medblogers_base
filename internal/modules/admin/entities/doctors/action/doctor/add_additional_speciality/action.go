package add_additional_speciality

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/add_additional_speciality/dal"
	commondal "medblogers_base/internal/modules/admin/entities/doctors/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetDoctorByID(ctx context.Context, doctorID int64) (*doctor.Doctor, error)
	GetSpecialityByID(ctx context.Context, cityID int64) (*speciality.Speciality, error)
}

type ActionDal interface {
	AddDoctorAdditionalSpeciality(ctx context.Context, doctorID, specialityID int64) error
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: commondal.NewRepository(pool),
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, doctorID, specialityID int64) error {
	_, err := a.commonDal.GetDoctorByID(ctx, doctorID)
	if err != nil {
		return err
	}

	_, err = a.commonDal.GetSpecialityByID(ctx, specialityID)
	if err != nil {
		return err
	}

	return a.actionDal.AddDoctorAdditionalSpeciality(ctx, doctorID, specialityID)
}
