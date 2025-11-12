package dao

import "medblogers_base/internal/modules/freelancers/domain/city"

type CityDAOWithDoctorID struct {
	CityDAO
	DoctorID int64 `db:"doctor_id"`
}

func (c CityDAOWithDoctorID) ToDomain() *city.City {
	return city.BuildCity(
		city.WithID(c.ID),
		city.WithName(c.Name),
	)
}
