package network

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_social_networks/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/social_network"
)

type Dal interface {
	GetNetworks(ctx context.Context, freelancerID int64) ([]*social_network.SocialNetwork, error)
}

type Service struct {
	dal Dal
}

func New(dal Dal) *Service {
	return &Service{
		dal: dal,
	}
}

func (s *Service) GetNetworks(ctx context.Context, freelancerID int64) ([]dto.Network, error) {
	networks, err := s.dal.GetNetworks(ctx, freelancerID)
	if err != nil {
		return nil, err
	}

	return lo.Map(networks, func(item *social_network.SocialNetwork, _ int) dto.Network {
		return dto.Network{
			ID:   item.ID(),
			Name: item.Name(),
			Slug: item.Slug(),
		}
	}), nil
}
