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

// GetAdditionalCities получение доп городов доктора
func (r *Repository) GetAdditionalCities(ctx context.Context, doctorID int64) ([]*city.City, error) {
	sql := `
		select c.id as "id", c.name as "name"
		from docstar_site_doctor_additional_cities adc
			join docstar_site_city c on adc.city_id = c.id
		where doctor_id = $1
	`

	var cities dao.CitiesDAO
	err := pgxscan.Select(ctx, r.db, &cities, sql, doctorID)
	if err != nil {
		return nil, err
	}
	return cities.ToDomain(), err
}

// GetAdditionalSpecialities получение доп специальностей доктора
func (r *Repository) GetAdditionalSpecialities(ctx context.Context, doctorID int64) ([]*speciality.Speciality, error) {
	sql := `
		select s.id as "id", s.name as "name"
		from docstar_site_doctor_additional_specialties ads
			join docstar_site_speciallity s on ads.speciallity_id = s.id
		where doctor_id = $1
	`

	var specialities dao.SpecialitiesDAO
	err := pgxscan.Select(ctx, r.db, &specialities, sql, doctorID)
	if err != nil {
		return nil, err
	}
	return specialities.ToDomain(), err
}
