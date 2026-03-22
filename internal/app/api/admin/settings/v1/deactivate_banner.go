package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) DeactivateBanner(ctx context.Context, req *desc.DeactivateBannerRequest) (resp *desc.DeactivateBannerResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/banner/{id}/deactivate", func(ctx context.Context) error {
		if err := i.admin.Actions.BannerModule.Deactivate.Do(ctx, req.GetBannerId()); err != nil {
			return err
		}

		resp = &desc.DeactivateBannerResponse{}
		return nil
	})
}
