package get_social_networks

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_social_networks/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_social_networks/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_social_networks/service/network"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	network *network.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		network: network.New(dal.NewRepository(pool)),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID int64) ([]dto.Network, error) {
	return a.network.GetNetworks(ctx, freelancerID)
}
