package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/v1"
)

func (i *Implementation) ActivateDoctor(ctx context.Context, req *desc.ActivateDoctorRequest) (resp *desc.ActivateDoctorResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/activate", func(ctx context.Context) error {
		return i.admin.Actions.DoctorModule.DoctorAgg.ActivateDoctor.Do(ctx, req.GetDoctorId())
	})
}
