package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) ChangeDoctorVipActivity(ctx context.Context, req *desc.ChangeDoctorVipActivityRequest) (resp *desc.ChangeDoctorVipActivityResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{doctor_id}/change_vip_activity", func(ctx context.Context) error {
		err := i.admin.Actions.DoctorModule.DoctorAgg.ChangeDoctorVipActivity.Do(ctx, req.GetDoctorId(), req.GetIsVipActive())
		if err != nil {
			return err
		}

		return nil
	})
}
