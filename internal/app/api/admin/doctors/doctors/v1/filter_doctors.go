package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
	"medblogers_base/internal/pkg/formatters"
)

func (i *Implementation) FilterDoctors(ctx context.Context, req *desc.FilterDoctorsRequest) (resp *desc.FilterDoctorsResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors/filter", func(ctx context.Context) error {

		res, err := i.admin.Actions.DoctorModule.DoctorAgg.FilterDoctors.Do(ctx, req.GetSpecialities())
		if err != nil {
			return err
		}

		resp = &desc.FilterDoctorsResponse{
			Doctors: lo.Map(res, func(item *doctor.Doctor, index int) *desc.FilterDoctorsResponse_Doctor {
				cooperationType := item.GetCooperationType()

				return &desc.FilterDoctorsResponse_Doctor{
					Id:       int64(item.GetID()),
					Name:     item.GetName(),
					IsActive: item.GetIsActive(),
					// todo Image
					CooperationType: &desc.CooperationType{
						Id:   cooperationType.ID(),
						Name: cooperationType.Name(),
					},
					CreatedAt: formatters.TimeRuFormat(item.GetCreatedAt()),
				}
			}),
		}
		return nil
	})
}
