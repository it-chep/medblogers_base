package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelansers/v1"
)

func (i *Implementation) SearchFreelancers(ctx context.Context, req *desc.SearchFreelancersRequest) (resp *desc.SearchFreelancersResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/deactivate", func(ctx context.Context) error {
		return nil
	})
}
