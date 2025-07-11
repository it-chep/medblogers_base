package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	cityDAO "github.com/it-chep/medblogers_base/internal/modules/doctors/dal/city_dal/dao"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/city"
)

type Repository struct {
}

// NewRepository создает новый репозиторий по работе с городами
func NewRepository() *Repository {
	return &Repository{}
}

// GetAllCities все города
func (r Repository) GetAllCities(ctx context.Context) ([]*city.City, error) {
	sql := `
		select c.id   as city_id,
			   c.name as city_name
		from docstar_site_city c
		group by c.id, c.name
		order by c.name
	`

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
