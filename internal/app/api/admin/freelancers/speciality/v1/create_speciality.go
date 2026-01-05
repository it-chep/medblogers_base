package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/speciality/v1"
)

func (i *Implementation) CreateSpeciality(ctx context.Context, req *desc.CreateSpecialityRequest) (resp *desc.CreateSpecialityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancers/specialities/create", func(ctx context.Context) error {
		return i.admin.Actions.FreelancerModule.SpecialityAgg.CreateSpeciality.Do(ctx, req.GetName())
	})
}
