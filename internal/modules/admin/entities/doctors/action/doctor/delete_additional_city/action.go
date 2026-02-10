package delete_additional_city

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/delete_additional_city/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	DeleteDoctorAdditionalCity(ctx context.Context, doctorID, cityID int64) error
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, doctorID, cityID int64) error {
	return a.actionDal.DeleteDoctorAdditionalCity(ctx, doctorID, cityID)
}
