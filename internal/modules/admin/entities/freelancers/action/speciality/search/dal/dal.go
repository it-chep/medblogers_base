package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SearchSpecialities(ctx context.Context, query string) (speciality.Specialities, error) {
	sql := `
		select id, name from freelancers_speciality where name ilike $1
		`

	var specialities dao.SpecialitiesDAO
	err := pgxscan.Select(ctx, r.db, &specialities, sql, query)
	if err != nil {
		return nil, err
	}
	return specialities.ToDomain(), err
}
