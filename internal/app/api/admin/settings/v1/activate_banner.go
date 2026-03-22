package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) ActivateBanner(ctx context.Context, req *desc.ActivateBannerRequest) (resp *desc.ActivateBannerResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/banner/{id}/activate", func(ctx context.Context) error {
		if err := i.admin.Actions.BannerModule.Activate.Do(ctx, req.GetBannerId()); err != nil {
			return err
		}

		resp = &desc.ActivateBannerResponse{}
		return nil
	})
}
