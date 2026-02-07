package delete

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/speciality/delete/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	DeleteSpeciality(ctx context.Context, id int64) error
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

func (a *Action) Do(ctx context.Context, id int64) error {
	return a.actionDal.DeleteSpeciality(ctx, id)
}
