package dal

import (
	"context"
	"fmt"
	"medblogers_base/internal/pkg/postgres"

	cityDAO "medblogers_base/internal/modules/doctors/dal/city_dal/dao"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	specialityDAO "medblogers_base/internal/modules/doctors/dal/speciality_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"

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

// SearchDoctors поиск докторов
func (r Repository) SearchDoctors(ctx context.Context, query string) ([]*doctor.Doctor, error) {
	sql := `
		select 
			name, slug
		from docstar_site_doctor d
		where d.is_active = true and d.name ilike $1
		order by d.name
		limit $2;
	`
	query = fmt.Sprintf("%s%s%s", "%", query, "%")
	var doctors []*dao.DoctorDAO
	// todo сделать лимит нормально
	if err := pgxscan.Select(ctx, r.db, &doctors, sql, query, 30); err != nil {
		return []*doctor.Doctor{}, err
	}

	result := make([]*doctor.Doctor, 0, len(doctors))
	for _, doctorDAO := range doctors {
		result = append(result, doctorDAO.ToDomain())
	}

	return result, nil
}

// SearchCities поиск городов
func (r Repository) SearchCities(ctx context.Context, query string) ([]*city.City, error) {
	sql := `
		select c.id 					 as city_id,
			   c.name                    as city_name,
			   count(distinct doctor_id) as doctors_count
		from docstar_site_city c
				 left join (select dc.city_id, dc.doctor_id
							from docstar_site_doctor_additional_cities dc
									 join docstar_site_doctor d on dc.doctor_id = d.id
							where d.is_active = true) as combined on c.id = combined.city_id
		where c.name ilike $1                   
		group by c.id, c.name
		having count(distinct doctor_id) != 0
		order by c.name
		limit $2;
	`
	query = fmt.Sprintf("%s%s%s", "%", query, "%")

	var cities []*cityDAO.CityDAOWithDoctorsCount
	// todo сделать лимит нормально
	if err := pgxscan.Select(ctx, r.db, &cities, sql, query, 5); err != nil {
		return []*city.City{}, err
	}

	result := make([]*city.City, 0, len(cities))
	for _, c := range cities {
		result = append(result, c.ToDomain())
	}

	return result, nil
}

// SearchSpecialities поиск специальностей
func (r Repository) SearchSpecialities(ctx context.Context, query string) ([]*speciality.Speciality, error) {
	sql := `
		select s.id                      as speciality_id,
			   s.name                    as speciality_name,
			   count(distinct doctor_id) as doctors_count
		from docstar_site_speciallity s
				 left join (select dc.speciallity_id, dc.doctor_id
							from docstar_site_doctor_additional_specialties dc
									 join docstar_site_doctor d on dc.doctor_id = d.id
							where d.is_active = true) as combined on s.id = combined.speciallity_id
		where s.name ilike $1                   
		group by s.id, s.name
		having count(distinct doctor_id) != 0
		order by s.name
		limit $2;
	`
	query = fmt.Sprintf("%s%s%s", "%", query, "%")

	var specialities []*specialityDAO.SpecialityDAOWithDoctorsCount
	// todo сделать лимит нормально
	if err := pgxscan.Select(ctx, r.db, &specialities, sql, query, 5); err != nil {
		return []*speciality.Speciality{}, err
	}

	result := make([]*speciality.Speciality, 0, len(specialities))
	for _, c := range specialities {
		result = append(result, c.ToDomain())
	}

	return result, nil
}
