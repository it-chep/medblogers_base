package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/city/v1"
)

func (i *Implementation) CreateCity(ctx context.Context, req *desc.CreateCityRequest) (resp *desc.CreateCityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancers/cities/create", func(ctx context.Context) error {
		return i.admin.Actions.FreelancerModule.CityAgg.CreateCity.Do(ctx, req.GetName())
	})
}
