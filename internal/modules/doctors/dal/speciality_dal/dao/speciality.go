package dao

import "medblogers_base/internal/modules/doctors/domain/speciality"

type SpecialityDAO struct {
	ID                   int64    `db:"id" json:"id"`
	Name                 string   `db:"name" json:"name"`
	IsOnlyAdditional     *bool    `db:"is_only_additional" json:"is_only_additional"`
	PrimarySpecialityIDs *[]int64 `db:"primary_speciality_ids" json:"primary_speciality_ids"`
}

func (s SpecialityDAO) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(
		speciality.WithID(speciality.SpecialityID(s.ID)),
		speciality.WithName(s.Name),
		speciality.WithIsOnlyAdditional(s.IsOnlyAdditional),
		speciality.WithPrimaryIDs(*s.PrimarySpecialityIDs),
	)
}
