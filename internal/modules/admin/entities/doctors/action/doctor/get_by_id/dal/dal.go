package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/doctors/dal/dao"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
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

// GetCity получение города доктора
func (r *Repository) GetCity(ctx context.Context, cityID int64) (*city.City, error) {
	sql := `
		select c.id as "id", c.name as "name"
		from docstar_site_city c
		where id = $1
	`

	var city dao.CityDAO
	err := pgxscan.Select(ctx, r.db, &city, sql, cityID)
	if err != nil {
		return nil, err
	}
	return city.ToDomain(), err
}

// GetSpeciality получение специальности доктора
func (r *Repository) GetSpeciality(ctx context.Context, specialityID int64) (*speciality.Speciality, error) {
	sql := `
		select s.id as "id", s.name as "name"
		from docstar_site_speciallity s
		where id = $1
	`

	var spec dao.SpecialityDAO
	err := pgxscan.Select(ctx, r.db, &spec, sql, specialityID)
	if err != nil {
		return nil, err
	}
	return spec.ToDomain(), err
}
