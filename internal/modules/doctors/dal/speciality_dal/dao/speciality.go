package dao

import "medblogers_base/internal/modules/doctors/domain/speciality"

type SpecialityDAO struct {
	ID                  int64  `db:"id" json:"id"`
	Name                string `db:"name" json:"name"`
	PrimarySpecialityID *int64 `db:"primary_speciality_id" json:"primary_speciality_id"`
	IsOnlyAdditional    *bool  `db:"is_only_additional" json:"is_only_additional"`
}

func (s SpecialityDAO) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(
		speciality.WithID(speciality.SpecialityID(s.ID)),
		speciality.WithName(s.Name),
		speciality.WithPrimarySpecialityID(s.PrimarySpecialityID),
		speciality.WithIsOnlyAdditional(s.IsOnlyAdditional),
	)
}
