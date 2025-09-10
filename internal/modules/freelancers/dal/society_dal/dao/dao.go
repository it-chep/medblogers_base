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

type Networks []SocialNetworkDao

func (n Networks) ToDomain() []*social_network.SocialNetwork {
	domain := make([]*social_network.SocialNetwork, 0, len(n))
	for _, network := range n {
		domain = append(domain, network.ToDomain())
	}
	return domain
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

type SocialNetworkWithFreelancerID struct {
	SocialNetworkDao
	FreelancerID int64 `db:"freelancer_id" json:"freelancer_id"`
}
