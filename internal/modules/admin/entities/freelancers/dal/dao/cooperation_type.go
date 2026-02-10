package dao

import (
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
)

type CooperationTypeDAO struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type CooperationTypes []CooperationTypeDAO

func (c *CooperationTypeDAO) ToDomain() *freelancer.CooperationType {
	return freelancer.NewCooperationType(c.ID, c.Name)
}

func (c CooperationTypes) ToDomain() []*freelancer.CooperationType {
	return lo.Map(c, func(item CooperationTypeDAO, index int) *freelancer.CooperationType {
		return item.ToDomain()
	})
}
