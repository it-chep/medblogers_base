package dal

import (
	"context"
	cityDAO "medblogers_base/internal/modules/doctors/dal/city_dal/dao"
	specialityDAO "medblogers_base/internal/modules/doctors/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"medblogers_base/internal/modules/doctors/domain/doctor"

	"github.com/georgysavva/scany/pgxscan"
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

// GetDoctorAdditionalCities получение информации о городах доктора
func (r Repository) GetDoctorAdditionalCities(ctx context.Context, doctorID doctor.MedblogersID) ([]*city.City, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных городов доктора")
	sql := `
		select c.id, c.name
        from docstar_site_city c
        	inner join docstar_site_doctor_additional_cities dc ON c.id = dc.city_id
        where dc.doctor_id = $1
        order by dc.id
	`

	var cities []*cityDAO.CityDAO
	if err := pgxscan.Select(ctx, r.db, &cities, sql, doctorID); err != nil {
		return nil, err
	}

	result := make([]*city.City, len(cities), 0)
	for _, c := range cities {
		result = append(result, c.ToDomain())
	}

	return result, nil
}

// GetDoctorAdditionalSpecialities получение информации о специальностях доктора
func (r Repository) GetDoctorAdditionalSpecialities(ctx context.Context, doctorID doctor.MedblogersID) ([]*speciality.Speciality, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных специальностей доктора")
	sql := `
		select s.id, s.name
		from docstar_site_speciallity s
			inner join docstar_site_doctor_additional_specialties dc ON s.id = dc.speciallity_id
		where dc.doctor_id = $1
        order by dc.id
	`

	var specialities []*specialityDAO.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db, &specialities, sql, doctorID); err != nil {
		return nil, err
	}

	result := make([]*speciality.Speciality, len(specialities), 0)
	for _, s := range specialities {
		result = append(result, s.ToDomain())
	}

	return result, nil
}
