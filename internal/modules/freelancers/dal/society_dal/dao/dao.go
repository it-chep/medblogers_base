package dao

import (
	"medblogers_base/internal/modules/freelancers/domain/social_network"
)

type SocialNetworkDao struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s SocialNetworkDao) ToDomain() *social_network.SocialNetwork {
	return social_network.BuildSocialNetwork(
		social_network.WithID(s.ID),
		social_network.WithName(s.Name),
	)
}

type SocialNetworkWithFreelancersCount struct {
	SocialNetworkDao
	FreelancersCount int64 `db:"freelancers_count" json:"freelancers_count"`
}

func (s SocialNetworkWithFreelancersCount) ToDomain() *social_network.SocialNetwork {
	return social_network.BuildSocialNetwork(
		social_network.WithID(s.ID),
		social_network.WithName(s.Name),
		social_network.WithFreelancersCount(s.FreelancersCount),
	)
}
