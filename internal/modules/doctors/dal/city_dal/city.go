package city_dal

import (
	"context"
	"fmt"
	cityDAO "medblogers_base/internal/modules/doctors/dal/city_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/city"
)

// Repository города
type Repository struct {
}

// NewRepository создает новый репозиторий по работе с городами
func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) GetCities(ctx context.Context) ([]*city.City, error) {
	sql := `
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
