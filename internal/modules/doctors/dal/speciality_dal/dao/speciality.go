package dao

import (
	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/speciality"
)

type SpecialityDAO struct {
	ID           int64  `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	DoctorsCount int64  `db:"doctors_count" json:"doctors_count"`
}

func (s SpecialityDAO) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(
		speciality.WithID(speciality.SpecialityID(s.ID)),
		speciality.WithName(s.Name),
		speciality.WithDoctorsCount(s.DoctorsCount),
	)
}
