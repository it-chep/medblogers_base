package dao

import (
	"medblogers_base/internal/modules/freelancers/domain/speciality"
)

type SpecialityDAOWithDoctorID struct {
	SpecialityDAO
	DoctorID int64 `db:"doctor_id"`
}

func (s SpecialityDAOWithDoctorID) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(
		speciality.WithID(s.ID),
		speciality.WithName(s.Name),
	)
}
