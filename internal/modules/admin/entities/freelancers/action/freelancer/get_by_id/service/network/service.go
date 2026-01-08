package network

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
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

func (s *Service) Enrich(ctx context.Context, freelancerDTO *dto.FreelancerDTO) (*dto.FreelancerDTO, error) {
	networks, err := s.dal.GetNetworks(ctx, freelancerDTO.ID)
	if err != nil {
		return nil, err
	}

	freelancerDTO.SocialNetworks = lo.Map(networks, func(item *social_network.SocialNetwork, _ int) dto.Network {
		return dto.Network{
			ID:   item.ID(),
			Name: item.Name(),
			Slug: item.Slug(),
		}
	})

	return freelancerDTO, nil
}
