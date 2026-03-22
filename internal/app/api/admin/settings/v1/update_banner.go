package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) UpdateBanner(ctx context.Context, req *desc.UpdateBannerRequest) (resp *desc.UpdateBannerResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/banner/{id}/update", func(ctx context.Context) error {
		err := i.admin.Actions.BannerModule.Update.Do(ctx, req.GetBannerId(), newUpdateRequest(
			req.GetName(),
			req.GetOrderingNumber(),
			req.GetBannerLink(),
		))
		if err != nil {
			return err
		}

		resp = &desc.UpdateBannerResponse{}
		return nil
	})
}
