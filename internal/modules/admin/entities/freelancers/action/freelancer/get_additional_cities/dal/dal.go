package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
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

// GetAdditionalCities получение только дополнительных городов
func (r *Repository) GetAdditionalCities(ctx context.Context, freelancerID int64) ([]*city.City, error) {
	sql := `
		select c.id, c.name 
		from freelancer_city_m2m m2m 
		    join freelancers_city c on m2m.city_id = c.id
         	join freelancer f on m2m.freelancer_id = f.id
		where m2m.freelancer_id = $1 and f.city_id != m2m.city_id
	`

	var cities dao.CitiesDAO
	err := pgxscan.Select(ctx, r.db, &cities, sql, freelancerID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return cities.ToDomain(), nil
}
