package dao

import "medblogers_base/internal/modules/doctors/domain/speciality"

type SpecialityDAO struct {
	ID                   int64    `db:"id" json:"id"`
	Name                 string   `db:"name" json:"name"`
	IsOnlyAdditional     *bool    `db:"is_only_additional" json:"is_only_additional"`
	PrimarySpecialityIDs *[]int64 `db:"primary_specialities_ids" json:"primary_specialities_ids"`
}

func (s SpecialityDAO) ToDomain() *speciality.Speciality {
	opts := []speciality.Option{
		speciality.WithID(speciality.SpecialityID(s.ID)),
		speciality.WithName(s.Name),
		speciality.WithIsOnlyAdditional(s.IsOnlyAdditional),
	}

	if s.PrimarySpecialityIDs != nil {
		opts = append(opts, speciality.WithPrimaryIDs(*s.PrimarySpecialityIDs))
	}

	return speciality.BuildSpeciality(opts...)
}
