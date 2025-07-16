package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	cityDAO "medblogers_base/internal/modules/doctors/dal/city_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/city"

	"github.com/georgysavva/scany/pgxscan"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с городами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// GetAllCities все города
func (r Repository) GetAllCities(ctx context.Context) ([]*city.City, error) {
	sql := `
		select c.id   as id,
			   c.name as name
		from docstar_site_city c
		group by c.id, c.name
		order by c.name
	`

	var citiesDAO []cityDAO.CityDAO
	if err := pgxscan.Select(ctx, r.db, &citiesDAO, sql); err != nil {
		return nil, err
	}

	cities := make([]*city.City, 0, len(citiesDAO))
	for _, dao := range citiesDAO {
		cities = append(cities, dao.ToDomain())
	}

	return cities, nil
}
