package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) AddNetwork(ctx context.Context, req *desc.AddNetworkRequest) (resp *desc.AddNetworkResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/add_network", func(ctx context.Context) error {
		return i.admin.Actions.FreelancerModule.FreelancerAgg.AddNetwork.Do(ctx, req.GetFreelancerId(), req.GetSocietyId())
	})
}
