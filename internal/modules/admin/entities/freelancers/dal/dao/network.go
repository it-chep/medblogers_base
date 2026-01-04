package dao

import (
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/social_network"
)

type NetworkDAO struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type NetworksDAO []NetworkDAO

func (n NetworkDAO) ToDomain() *social_network.SocialNetwork {
	return social_network.BuildSocialNetwork(social_network.WithID(n.ID), social_network.WithName(n.Name))
}

func (n NetworksDAO) ToDomain() []*social_network.SocialNetwork {
	return lo.Map(n, func(item NetworkDAO, _ int) *social_network.SocialNetwork {
		return item.ToDomain()
	})
}
