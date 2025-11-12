package dao

import "medblogers_base/internal/modules/doctors/domain/city"

type CityDAO struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (c CityDAO) ToDomain() *city.City {
	return city.BuildCity(
		city.WithID(city.CityID(c.ID)),
		city.WithName(c.Name),
	)
}
