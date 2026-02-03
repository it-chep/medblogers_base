package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) GetDoctorCooperationTypes(ctx context.Context, _ *desc.GetDoctorCooperationTypesRequest) (resp *desc.GetDoctorCooperationTypesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/cooperation_types", func(ctx context.Context) error {
		coopTypes, err := i.admin.Actions.DoctorModule.DoctorAgg.GetCooperationTypes.Do(ctx)
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorCooperationTypesResponse{
			CooperationTypes: lo.Map(coopTypes, func(item *doctor.CooperationType, index int) *desc.CooperationType {
				return &desc.CooperationType{
					Id:   item.ID(),
					Name: item.Name(),
				}
			}),
		}
		return nil
	})
}
