package search

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/speciality/search/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	SearchSpecialities(ctx context.Context, name string) ([]*speciality.Speciality, error)
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

func (a *Action) Do(ctx context.Context, name string) ([]*speciality.Speciality, error) {
	return a.actionDal.SearchSpecialities(ctx, name)
}
