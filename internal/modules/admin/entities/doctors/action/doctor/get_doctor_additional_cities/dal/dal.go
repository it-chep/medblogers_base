package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
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

// GetAdditionalCities получение доп городов доктора
func (r *Repository) GetAdditionalCities(ctx context.Context, doctorID int64) ([]*city.City, error) {
	sql := `
		select c.id as "id", c.name as "name"
		from docstar_site_doctor_additional_cities adc
			join docstar_site_city c on adc.city_id = c.id
			join docstar_site_doctor d on adc.doctor_id = d.id
		where adc.doctor_id = $1 and adc.city_id != d.city_id
	`

	var cities dao.CitiesDAO
	err := pgxscan.Select(ctx, r.db, &cities, sql, doctorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return cities.ToDomain(), err
}
