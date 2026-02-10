package search

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/speciality/search/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	SearchSpecialities(ctx context.Context, query string) (speciality.Specialities, error)
}
type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, query string) (speciality.Specialities, error) {
	return a.actionDal.SearchSpecialities(ctx, query)
}
