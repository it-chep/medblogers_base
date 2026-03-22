package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) CreateBanner(ctx context.Context, req *desc.CreateBannerRequest) (resp *desc.CreateBannerResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/banner/create", func(ctx context.Context) error {
		bannerID, err := i.admin.Actions.BannerModule.Create.Do(ctx, newUpdateRequest(
			req.GetName(),
			req.GetOrderingNumber(),
			req.GetBannerLink(),
		))
		if err != nil {
			return err
		}

		resp = &desc.CreateBannerResponse{
			BannerId: bannerID,
		}

		return nil
	})
}
