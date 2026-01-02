package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
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

func (r *Repository) GetCities(ctx context.Context) ([]*city.City, error) {
	sql := `
		select id, name from freelancers_city
		order by name asc
		`

	var cities dao.CitiesDAO
	err := pgxscan.Select(ctx, r.db, &cities, sql)
	if err != nil {
		return nil, err
	}
	return cities.ToDomain(), nil
}
