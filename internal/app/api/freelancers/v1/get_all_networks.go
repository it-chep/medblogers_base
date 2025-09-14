package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
)

func (i *Implementation) GetSocialNetworks(ctx context.Context, _ *desc.GetSocialNetworksRequest) (*desc.GetSocialNetworksResponse, error) {
	networks, err := i.freelancers.Actions.GetAllNetworks.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetSocialNetworksResponse{
		SocialNetworks: lo.Map(networks, func(item *social_network.SocialNetwork, _ int) *desc.GetSocialNetworksResponse_SocialNetworkItem {
			return &desc.GetSocialNetworksResponse_SocialNetworkItem{
				Id:   item.ID(),
				Name: item.Name(),
			}
		}),
	}, nil
}
