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
func (r Repository) GetDoctorAdditionalCities(ctx context.Context, doctorID doctor.MedblogersID) (map[city.CityID]*city.City, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных городов доктора")
	sql := `
		select c.id, c.name
        from docstar_site_city c
        inner join docstar_site_doctor_additional_cities dc ON c.id = dc.city_id
        where dc.doctor_id = $1
        order by c.name
	`

	var cities []*cityDAO.CityDAO
	if err := pgxscan.Select(ctx, r.db, &cities, sql, doctorID); err != nil {
		return nil, err
	}

	result := make(map[city.CityID]*city.City, len(cities))
	for _, c := range cities {
		result[city.CityID(c.ID)] = c.ToDomain()
	}

	return result, nil
}

// GetDoctorAdditionalSpecialities получение информации о специальностях доктора
func (r Repository) GetDoctorAdditionalSpecialities(ctx context.Context, doctorID doctor.MedblogersID) (map[speciality.SpecialityID]*speciality.Speciality, error) {
	logger.Message(ctx, "[Dal] Получение дополнительных специальностей доктора")
	sql := `
		select s.id, s.name
		from docstar_site_speciallity s
		inner join docstar_site_doctor_additional_specialties dc ON s.id = dc.speciallity_id
		where dc.doctor_id = $1
        order by s.name
	`

	var specialities []*specialityDAO.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db, &specialities, sql, doctorID); err != nil {
		return nil, err
	}

	result := make(map[speciality.SpecialityID]*speciality.Speciality, len(specialities))
	for _, s := range specialities {
		result[speciality.SpecialityID(s.ID)] = s.ToDomain()
	}

	return result, nil
}
