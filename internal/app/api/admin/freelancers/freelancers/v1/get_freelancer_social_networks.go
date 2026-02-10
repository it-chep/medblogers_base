package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_social_networks/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancerSocialNetworks(ctx context.Context, req *desc.GetFreelancerSocialNetworksRequest) (resp *desc.GetFreelancerSocialNetworksResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/social_networks", func(ctx context.Context) error {
		networks, err := i.admin.Actions.FreelancerModule.FreelancerAgg.GetNetworks.Do(ctx, req.GetFreelancerId())
		if err != nil {
			return err
		}

		resp = &desc.GetFreelancerSocialNetworksResponse{
			SocialNetworks: lo.Map(networks, func(item dto.Network, index int) *desc.Society {
				return &desc.Society{
					Id:   item.ID,
					Name: item.Name,
				}
			}),
		}
		return nil
	})
}
