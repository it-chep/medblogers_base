package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/doctors/dal/dao"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
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

func (r *Repository) SearchSpecialities(ctx context.Context, name string) ([]*speciality.Speciality, error) {
	sql := `
		select id, name from docstar_site_speciallity where name ilike $1
		order by name asc
		`

	var specialities dao.SpecialitiesDAO
	err := pgxscan.Select(ctx, r.db, &specialities, sql, name)
	if err != nil {
		return nil, err
	}
	return specialities.ToDomain(), nil
}
