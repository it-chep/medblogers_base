package social_network

import (
	"medblogers_base/internal/modules/admin/entities/freelancers/action/social_network/get"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/social_network/search"
	"medblogers_base/internal/pkg/postgres"
)

type FreelancerNetworkAggregator struct {
	GetNetworks    *get.Action
	SearchNetworks *search.Action
}

func New(pool postgres.PoolWrapper) *FreelancerNetworkAggregator {
	return &FreelancerNetworkAggregator{
		GetNetworks:    get.New(pool),
		SearchNetworks: search.New(pool),
	}
}
