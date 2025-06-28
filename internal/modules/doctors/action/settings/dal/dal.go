package dal

import (
	"context"
	"fmt"
	cityDAO "medblogers_base/internal/modules/doctors/dal/city_dal/dao"
	specialityDAO "medblogers_base/internal/modules/doctors/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

type Repository struct {
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) GetSpecialities(ctx context.Context) ([]*speciality.Speciality, error) {
	sql := fmt.Sprintf(`
		select s.id                      as speciality_id,
			   s.name                    as speciality_name,
			   count(distinct doctor_id) as doctors_count
		from docstar_site_speciallity s
				 left join (select dc.speciallity_id, dc.doctor_id
							from docstar_site_doctor_additional_specialties dc
									 join docstar_site_doctor d on dc.doctor_id = d.id
							where d.is_active = true) as combined on s.id = combined.speciallity_id
		group by s.id, s.name
		order by s.name
	`)

	var specialitiesDAO []specialityDAO.SpecialityDAO
	if err := pgxscan.Select(ctx, r.db.Pool(ctx), &specialitiesDAO, sql); err != nil {
		return nil, err
	}

	specialities := make([]*speciality.Speciality, 0, len(specialitiesDAO))
	for _, dao := range specialitiesDAO {
		specialities = append(specialities, dao.ToDomain())
	}

	return specialities, nil
}

func (r Repository) GetCities(ctx context.Context) ([]*city.City, error) {
	sql := fmt.Sprintf(`
		select c.id                      as city_id,
			   c.name                    as city_name,
			   count(distinct doctor_id) as doctors_count
		from docstar_site_city c
				 left join (select dc.city_id, dc.doctor_id
							from docstar_site_doctor_additional_cities dc
									 join docstar_site_doctor d on dc.doctor_id = d.id
							where d.is_active = true) as combined on c.id = combined.city_id
		group by c.id, c.name
		order by c.name
	`)

	var citiesDAO []cityDAO.CityDAO
	if err := pgxscan.Select(ctx, r.db.Pool(ctx), &citiesDAO, sql); err != nil {
		return nil, err
	}

	cities := make([]*city.City, 0, len(citiesDAO))
	for _, dao := range citiesDAO {
		cities = append(cities, dao.ToDomain())
	}
	return cities, nil
}

func (r Repository) GetDoctorsCount(ctx context.Context) (int64, error) {
	sql := fmt.Sprintf(`
		select count(*) as doctors_count
		from docstar_site_doctor d
		where d.is_active = true
	`)

	var count int64
	if err := pgxscan.Select(ctx, r.db.Pool(ctx), &count, sql); err != nil {
		return 0, err
	}

	return count, nil
}
