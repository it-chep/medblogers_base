package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) DeleteAdditionalSpeciality(ctx context.Context, req *desc.DeleteDoctorAdditionalSpecialityRequest) (resp *desc.DeleteDoctorAdditionalSpecialityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/delete_additional_speciality", func(ctx context.Context) error {
		return i.admin.Actions.DoctorModule.DoctorAgg.DeleteAdditionalSpeciality.Do(ctx, req.GetDoctorId(), req.GetSpecialityId())
	})
}
