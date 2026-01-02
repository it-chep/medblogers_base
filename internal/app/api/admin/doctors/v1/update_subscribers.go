package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update_subscribers/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/v1"
)

func (i *Implementation) UpdateSubscribers(ctx context.Context, req *desc.UpdateSubscribersRequest) (resp *desc.UpdateSubscribersResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/update_subscribers", func(ctx context.Context) error {
		return i.admin.Actions.DoctorModule.DoctorAgg.UpdateSubscribers.Do(ctx, req.GetDoctorId(), dto.UpdateSubscribersRequest{}) // todo поправить контракт
	})
}
