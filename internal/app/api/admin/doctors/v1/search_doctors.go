package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/v1"
)

func (i *Implementation) SearchDoctors(ctx context.Context, req *desc.SearchDoctorsRequest) (resp *desc.SearchDoctorsResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors/search", func(ctx context.Context) error {

		return nil
	})
}
