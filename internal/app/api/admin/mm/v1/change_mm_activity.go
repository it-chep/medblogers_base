package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/mastermind/v1"
)

func (i *Implementation) ChangeMMActivity(ctx context.Context, req *desc.ChangeMMActivityRequest) (resp *desc.ChangeMMActivityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/mm/{id}/change_activity", func(ctx context.Context) error {
		resp = &desc.ChangeMMActivityResponse{}

		err := i.admin.Actions.MMModule.ChangeMMActivity.Do(ctx, req.GetMmId(), req.GetActivity())
		if err != nil {
			return err
		}

		return nil
	})
}
