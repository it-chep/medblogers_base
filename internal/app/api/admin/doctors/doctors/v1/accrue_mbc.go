package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) DoctorAccrueMBC(ctx context.Context, req *desc.DoctorAccrueMBCRequest) (resp *desc.DoctorAccrueMBCResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{doctor_id}/accure_mbc", func(ctx context.Context) error {
		return i.admin.Actions.DoctorModule.DoctorAgg.AccrueMBC.Do(ctx, req.GetDoctorId(), req.GetMbcCount())
	})
}
