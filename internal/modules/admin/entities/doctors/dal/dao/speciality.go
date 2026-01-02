package dao

import (
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
)

type SpecialityDAO struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type SpecialitiesDAO []SpecialityDAO

func (s SpecialityDAO) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(speciality.WithID(speciality.SpecialityID(s.ID)), speciality.WithName(s.Name))
}

func (c SpecialitiesDAO) ToDomain() []*speciality.Speciality {
	return lo.Map(c, func(item SpecialityDAO, _ int) *speciality.Speciality {
		return item.ToDomain()
	})
}
