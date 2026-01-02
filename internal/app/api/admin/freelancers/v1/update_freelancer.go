package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelansers/v1"
)

func (i *Implementation) UpdateFreelancer(ctx context.Context, req *desc.UpdateFreelancerRequest) (resp *desc.UpdateFreelancerResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/deactivate", func(ctx context.Context) error {
		return nil
	})
}
