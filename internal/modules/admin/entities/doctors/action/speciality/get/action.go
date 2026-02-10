package get

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/speciality/get/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetSpecialities(ctx context.Context) ([]*speciality.Speciality, error)
}

// Action получение городов
type Action struct {
	actionDal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]*speciality.Speciality, error) {
	return a.actionDal.GetSpecialities(ctx)
}
