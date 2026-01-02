package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/doctors/dal/dao"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
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

func (r *Repository) SearchCities(ctx context.Context, name string) ([]*city.City, error) {
	sql := `
		select id, name from docstar_site_city where name ilike $1
		order by name asc
		`

	var cities dao.CitiesDAO
	err := pgxscan.Select(ctx, r.db, &cities, sql, name)
	if err != nil {
		return nil, err
	}
	return cities.ToDomain(), nil
}
