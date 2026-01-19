package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/mm/v1"
)

func (i *Implementation) ManualNotificationMM(ctx context.Context, req *desc.ManualNotificationMMRequest) (resp *desc.ManualNotificationMMResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/mm/{id}/manual_notification", func(ctx context.Context) error {
		resp = &desc.ManualNotificationMMResponse{}

		err := i.admin.Actions.MMModule.ManualNotificationMM.Do(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
