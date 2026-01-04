package get

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetSpecialities(ctx context.Context) (speciality.Specialities, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) (speciality.Specialities, error) {
	return a.actionDal.GetSpecialities(ctx)
}
