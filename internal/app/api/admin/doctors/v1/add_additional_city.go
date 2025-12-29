package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/v1"
)

func (i *Implementation) AddAdditionalCity(ctx context.Context, req *desc.AddDoctorAdditionalCityRequest) (resp *desc.AddDoctorAdditionalCityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/add_additional_city", func(ctx context.Context) error {
		return i.admin.Actions.DoctorModule.DoctorAgg.DoctorAgg.AddAdditionalCity.Do(ctx, req.GetDoctorId(), req.GetCityId())
	})
}
