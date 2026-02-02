package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) GetDoctorCooperationTypes(ctx context.Context, req *desc.GetDoctorCooperationTypesRequest) (resp *desc.GetDoctorCooperationTypesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/cooperation_types", func(ctx context.Context) error {

		_, err := i.admin.Actions.DoctorModule.DoctorAgg.GetDoctors.Do(ctx)
		if err != nil {
			return err
		}

		//resp = &desc.GetDoctorCooperationTypesResponse{
		//	CooperationTypes: lo.Map(res, func(item *doctor.Doctor, index int) *desc.CooperationType {
		//		return &desc.CooperationType{
		//			Id:   int64(item.GetID()),
		//			Name: item.GetName(),
		//		}
		//	}),
		//}
		return nil
	})
}
