package dao

import "medblogers_base/internal/modules/doctors/domain/city"

type CityDAOWithDoctorsCount struct {
	CityDAO
	DoctorsCount int64 `db:"doctors_count" json:"doctors_count"`
}

func (c CityDAOWithDoctorsCount) ToDomain() *city.City {
	return city.BuildCity(
		city.WithID(city.CityID(c.ID)),
		city.WithName(c.Name),
		city.WithDoctorsCount(c.DoctorsCount),
	)
}
