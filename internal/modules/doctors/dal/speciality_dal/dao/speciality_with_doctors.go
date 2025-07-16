package dao

import (
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

type SpecialityDAOWithDoctorsCount struct {
	SpecialityDAO
	DoctorsCount int64 `db:"doctors_count" json:"doctors_count"`
}

func (s SpecialityDAOWithDoctorsCount) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(
		speciality.WithID(speciality.SpecialityID(s.ID)),
		speciality.WithName(s.Name),
		speciality.WithDoctorsCount(s.DoctorsCount),
	)
}
