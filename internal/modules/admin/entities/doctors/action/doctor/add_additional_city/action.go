package add_additional_city

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/add_additional_city/dal"
	commondal "medblogers_base/internal/modules/admin/entities/doctors/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetDoctorByID(ctx context.Context, doctorID int64) (*doctor.Doctor, error)
	GetCityByID(ctx context.Context, cityID int64) (*city.City, error)
}

type ActionDal interface {
	AddDoctorAdditionalCity(ctx context.Context, doctorID, cityID int64) error
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

func (a *Action) Do(ctx context.Context, doctorID, cityID int64) error {
	_, err := a.commonDal.GetDoctorByID(ctx, doctorID)
	if err != nil {
		return err
	}

	_, err = a.commonDal.GetCityByID(ctx, cityID)
	if err != nil {
		return err
	}

	return a.actionDal.AddDoctorAdditionalCity(ctx, doctorID, cityID)
}
