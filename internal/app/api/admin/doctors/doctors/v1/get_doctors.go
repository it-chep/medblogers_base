package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) GetDoctors(ctx context.Context, req *desc.GetDoctorsRequest) (resp *desc.GetDoctorsResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors", func(ctx context.Context) error {

		res, err := i.admin.Actions.DoctorModule.DoctorAgg.GetDoctors.Do(ctx)
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorsResponse{
			Doctors: lo.Map(res, func(item *doctor.Doctor, index int) *desc.GetDoctorsResponse_Doctor {
				cooperationType := item.GetCooperationType()

				return &desc.GetDoctorsResponse_Doctor{
					Id:       int64(item.GetID()),
					Name:     item.GetName(),
					IsActive: item.GetIsActive(),
					// todo Image
					CooperationType: &desc.CooperationType{
						Id:   cooperationType.ID(),
						Name: cooperationType.Name(),
					},
				}
			}),
		}
		return nil
	})
}
