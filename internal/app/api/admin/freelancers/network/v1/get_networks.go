package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/network/v1"
)

func (i *Implementation) GetNetworks(ctx context.Context, req *desc.GetNetworksRequest) (resp *desc.GetNetworksResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancers/networks", func(ctx context.Context) error {
		_, err := i.admin.Actions.FreelancerModule.NetworksAgg.GetNetworks.Do(ctx)
		if err != nil {
			return err
		}
		// todo
		//resp = &desc.GetNetworksResponse{
		//	: lo.Map(res, func(item *city.City, _ int) *desc.SearchCitiesResponse_CityItem {
		//		return &desc.SearchCitiesResponse_CityItem{
		//			Id:   int64(item.ID()),
		//			Name: item.Name(),
		//		}
		//	}),
		//}

		return nil
	})
}
