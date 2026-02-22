package filter_doctors

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/filter_doctors/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"
)

type Repository interface {
	FilterDoctors(ctx context.Context, specialitiesIDs []int64) ([]*doctor.Doctor, error)
}

type Action struct {
	dal Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, specialitiesIDs []int64) ([]*doctor.Doctor, error) {
	return a.dal.FilterDoctors(ctx, specialitiesIDs)
}
