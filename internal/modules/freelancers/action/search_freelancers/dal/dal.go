package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"

	pkgconfig "medblogers_base/internal/pkg/config"
	"medblogers_base/internal/pkg/postgres"
)

const (
	defaultFreelancersLimit  int64 = 30
	defaultCitiesLimit       int64 = 5
	defaultSpecialitiesLimit int64 = 5
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

// SearchFreelancers поиск фрилансеров
func (r Repository) SearchFreelancers(ctx context.Context, query string) ([]*doctor.Doctor, error) {
	sql := `
		select 
			f.id, f.name, f.slug, f.price_category, f.is_worked_with_freelancers, c.name as "city_name", s.name as "speciality_name"
		from freelancer f
		join freelancers_city c on f.city_id = c.id
		join freelancers_speciality s on f.speciallity_id = s.id
		where f.is_active = true and f.name ilike $1
		order by f.name
		limit $2;
	`
	query = fmt.Sprintf("%s%s%s", "%", query, "%")
	limit := r.getLimit(ctx, defaultFreelancersLimit, config.SearchDoctorsLimit)

	var freelancers []*dao.freelancersearchDAO
	if err := pgxscan.Select(ctx, r.db, &freelancers, sql, query, limit); err != nil {
		return []*doctor.Doctor{}, err
	}

	result := make([]*doctor.Doctor, 0, len(freelancers))
	for _, doctorDAO := range freelancers {
		result = append(result, doctorDAO.ToDomain())
	}

	return result, nil
}

// SearchCities поиск городов
func (r Repository) SearchCities(ctx context.Context, query string) ([]*city.City, error) {
	sql := `
		select c.id 					 as id,
			   c.name                    as name,
			   count(distinct freelancer_id) as freelancers_count
		from freelancers_city c
				 left join (select fc.city_id, fc.freelancer_id
							from freelancer_city_m2m fc
									 join freelancer f on fc.freelancer_id = f.id
							where f.is_active = true) as combined on c.id = combined.city_id
		where c.name ilike $1                   
		group by c.id, c.name
		having count(distinct freelancer_id) != 0
		order by freelancers_count desc
		limit $2;
	`
	query = fmt.Sprintf("%s%s%s", "%", query, "%")
	limit := r.getLimit(ctx, defaultCitiesLimit, config.SearchCitiesLimit)

	var cities []*cityDAO.CityDAOWithfreelancersCount
	if err := pgxscan.Select(ctx, r.db, &cities, sql, query, limit); err != nil {
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
		select s.id                      as id,
			   s.name                    as name,
			   count(distinct freelancer_id) as freelancers_count
		from freelancers_speciality s
				 left join (select fs.speciallity_id, fs.freelancer_id
							from freelancer_speciality_m2m fs
									 join freelancer f on fs.freelancer_id = f.id
							where f.is_active = true) as combined on s.id = combined.speciality_id
		where s.name ilike $1                   
		group by s.id, s.name
		having count(distinct freelancer_id) != 0
		order by freelancers_count desc 
		limit $2;
	`
	query = fmt.Sprintf("%s%s%s", "%", query, "%")
	limit := r.getLimit(ctx, defaultSpecialitiesLimit, config.SearchSpecialitiesLimit)

	var specialities []*specialityDAO.SpecialityDAOWithfreelancersCount
	if err := pgxscan.Select(ctx, r.db, &specialities, sql, query, limit); err != nil {
		return []*speciality.Speciality{}, err
	}

	result := make([]*speciality.Speciality, 0, len(specialities))
	for _, c := range specialities {
		result = append(result, c.ToDomain())
	}

	return result, nil
}

// getLimit получение лимита поиска из конфига базы с фолбеком на дефолтные значения из кода
func (r Repository) getLimit(ctx context.Context, defaultVal int64, key pkgconfig.ConfigKey) int64 {
	// todo мб вынести это на уровень модуля и как-то красиво сделать
	limit := defaultVal
	val, _ := pkgconfig.GetValue(ctx, key)
	if val != nil {
		limit = val.Int64()
	}

	return limit
}
