package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/speciality/v1"
)

func (i *Implementation) GetSpecialities(ctx context.Context, req *desc.GetDoctorsSpecialitiesRequest) (resp *desc.GetDoctorsSpecialitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors/specialities", func(ctx context.Context) error {
		res, err := i.admin.Actions.DoctorModule.SpecialityAgg.GetSpecialities.Do(ctx)
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorsSpecialitiesResponse{
			Specialities: lo.Map(res, func(item *speciality.Speciality, _ int) *desc.GetDoctorsSpecialitiesResponse_SpecialityItem {
				return &desc.GetDoctorsSpecialitiesResponse_SpecialityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}

		return nil
	})
}
