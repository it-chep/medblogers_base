package create

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/speciality/create/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	CreateSpeciality(ctx context.Context, name string, isOnlyAdditional bool) error
}

// Action поиск городов
type Action struct {
	actionDal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, name string, isOnlyAdditional bool) error {
	return a.actionDal.CreateSpeciality(ctx, name, isOnlyAdditional)
}
