package dao

import (
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
)

type CooperationTypeDAO struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type CooperationTypes []CooperationTypeDAO

func (c *CooperationTypeDAO) ToDomain() *doctor.CooperationType {
	return doctor.NewCooperationType(c.ID, c.Name)
}

func (c CooperationTypes) ToDomain() []*doctor.CooperationType {
	return lo.Map(c, func(item CooperationTypeDAO, index int) *doctor.CooperationType {
		return item.ToDomain()
	})
}
