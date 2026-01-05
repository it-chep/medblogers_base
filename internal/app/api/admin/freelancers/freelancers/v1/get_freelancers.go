package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancers(ctx context.Context, req *desc.GetFreelancersRequest) (resp *desc.GetFreelancersResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/deactivate", func(ctx context.Context) error {
		return nil
	})
}
