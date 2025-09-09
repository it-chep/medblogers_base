package city_dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/freelancers/dal/city_dal/dao"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Repository города
type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с городами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// GetCitiesWithFreelancersCount получение списка городов с количеством фрилансеров
func (r Repository) GetCitiesWithFreelancersCount(ctx context.Context) ([]*city.City, error) {
	logger.Message(ctx, "[DAL] Запрос городов")
	sql := `
		with active_freelancers_in_cities as (
			select
				fc.city_id,
				fc.freelancer_id
			from freelancer_city_m2m fc
					 join freelancer f on fc.freelancer_id = f.id
			where f.is_active = true
		)
		select
			c.id as id,
			c.name as name,
			count(distinct af.freelancer_id) as freelancers_count
		from freelancers_city c
				 left join active_freelancers_in_cities af on c.id = af.city_id
		group by c.id, c.name
		having count(distinct af.freelancer_id) > 0
		order by c.name;
	`

	var citiesDAO []dao.CityDAOWithFreelancersCount
	if err := pgxscan.Select(ctx, r.db, &citiesDAO, sql); err != nil {
		return nil, err
	}

	cities := make([]*city.City, 0, len(citiesDAO))
	for _, cityDao := range citiesDAO {
		cities = append(cities, cityDao.ToDomain())
	}

	return cities, nil
}

// GetAllCities все города
func (r Repository) GetAllCities(ctx context.Context) ([]*city.City, error) {
	sql := `
		select c.id   as id,
			   c.name as name
		from freelancers_city c
		group by c.id, c.name
		order by c.name
	`

	var citiesDAO []dao.CityDAO
	if err := pgxscan.Select(ctx, r.db, &citiesDAO, sql); err != nil {
		return nil, err
	}

	cities := make([]*city.City, 0, len(citiesDAO))
	for _, cityDao := range citiesDAO {
		cities = append(cities, cityDao.ToDomain())
	}

	return cities, nil
}
