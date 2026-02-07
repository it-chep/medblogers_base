package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) GetDoctorsByIDs(ctx context.Context, req *desc.GetDoctorsByIDsRequest) (resp *desc.GetDoctorsByIDsResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors_by_ids", func(ctx context.Context) error {

		return nil
	})
}
