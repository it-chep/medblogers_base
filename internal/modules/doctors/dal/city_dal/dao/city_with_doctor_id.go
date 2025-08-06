package dao

import "medblogers_base/internal/modules/doctors/domain/city"

type CityDAOWithDoctorID struct {
	CityDAO
	DoctorID int64 `db:"doctor_id"`
}

func (c CityDAOWithDoctorID) ToDomain() *city.City {
	return city.BuildCity(
		city.WithID(city.CityID(c.ID)),
		city.WithName(c.Name),
	)
}
