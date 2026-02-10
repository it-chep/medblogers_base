package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) GetDoctorAdditionalSpecialities(ctx context.Context, req *desc.GetDoctorAdditionalSpecialitiesRequest) (resp *desc.GetDoctorAdditionalSpecialitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/additional_specialities", func(ctx context.Context) error {

		specialities, err := i.admin.Actions.DoctorModule.DoctorAgg.GetDoctorAdditionalSpecialities.Do(ctx, req.GetDoctorId())
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorAdditionalSpecialitiesResponse{
			AdditionalSpecialities: lo.Map(specialities, func(item *speciality.Speciality, index int) *desc.SpecialityItem {
				return &desc.SpecialityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}
		return nil
	})
}
