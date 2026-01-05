package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/city/v1"
)

func (i *Implementation) CreateCity(ctx context.Context, req *desc.CreateDoctorsCityRequest) (resp *desc.CreateDoctorsCityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors/search", func(ctx context.Context) error {
		return i.admin.Actions.DoctorModule.CityAgg.CreateCity.Do(ctx, req.GetName())
	})
}
