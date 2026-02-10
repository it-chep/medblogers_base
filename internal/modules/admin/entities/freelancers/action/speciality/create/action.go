package create

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/speciality/create/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	CreateSpeciality(ctx context.Context, name string) error
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, name string) error {
	return a.actionDal.CreateSpeciality(ctx, name)
}
