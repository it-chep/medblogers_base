package dao

import "medblogers_base/internal/modules/doctors/domain/speciality"

type SpecialityDAO struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s SpecialityDAO) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(
		speciality.WithID(speciality.SpecialityID(s.ID)),
		speciality.WithName(s.Name),
	)
}
