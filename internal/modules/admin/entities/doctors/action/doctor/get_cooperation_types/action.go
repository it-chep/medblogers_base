package get_cooperation_types

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_cooperation_types/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"
)

type Dal interface {
	GetDoctorCooperationTypes(ctx context.Context) ([]*doctor.CooperationType, error)
}

type Action struct {
	dal Dal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]*doctor.CooperationType, error) {
	return a.dal.GetDoctorCooperationTypes(ctx)
}
