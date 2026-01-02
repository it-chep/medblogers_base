package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/v1"
)

func (i *Implementation) SortDoctors(ctx context.Context, req *desc.SortDoctorsRequest) (resp *desc.SortDoctorsResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors/sort", func(ctx context.Context) error {

		return nil
	})
}
