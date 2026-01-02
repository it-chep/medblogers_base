package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/v1"
)

func (i *Implementation) AddAdditionalSpeciality(ctx context.Context, req *desc.AddDoctorAdditionalSpecialityRequest) (resp *desc.AddDoctorAdditionalSpecialityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/add_additional_speciality", func(ctx context.Context) error {
		return i.admin.Actions.DoctorModule.DoctorAgg.AddAdditionalSpeciality.Do(ctx, req.GetDoctorId(), req.GetSpecialityId())
	})
}
