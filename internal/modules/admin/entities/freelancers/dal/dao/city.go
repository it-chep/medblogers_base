package dao

import (
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
)

type CityDAO struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CitiesDAO []CityDAO

func (c CityDAO) ToDomain() *city.City {
	return city.BuildCity(city.WithID(city.CityID(c.ID)), city.WithName(c.Name))
}

func (c CitiesDAO) ToDomain() []*city.City {
	return lo.Map(c, func(item CityDAO, _ int) *city.City {
		return item.ToDomain()
	})
}
