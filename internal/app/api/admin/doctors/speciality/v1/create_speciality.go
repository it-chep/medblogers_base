package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/speciality/v1"
)

func (i *Implementation) CreateSpeciality(ctx context.Context, req *desc.CreateDoctorsSpecialityRequest) (resp *desc.CreateDoctorsSpecialityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors/specialities/create", func(ctx context.Context) error {
		err := i.admin.Actions.DoctorModule.SpecialityAgg.CreateSpeciality.Do(ctx, req.GetName(), false)
		if err != nil {
			return err
		}

		return nil
	})
}
