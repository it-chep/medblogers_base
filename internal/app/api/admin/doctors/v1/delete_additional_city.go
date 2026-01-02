package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/v1"
)

func (i *Implementation) DeleteAdditionalCity(ctx context.Context, req *desc.DeleteDoctorAdditionalCityRequest) (resp *desc.DeleteDoctorAdditionalCityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/delete_additional_city", func(ctx context.Context) error {
		return i.admin.Actions.DoctorModule.DoctorAgg.DeleteAdditionalCity.Do(ctx, req.GetDoctorId(), req.GetCityId())
	})
}
