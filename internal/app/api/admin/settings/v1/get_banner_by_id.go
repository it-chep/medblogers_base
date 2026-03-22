package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) GetBannerByID(ctx context.Context, req *desc.GetBannerByIDRequest) (resp *desc.GetBannerByIDResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/banner/{id}", func(ctx context.Context) error {
		banner, err := i.admin.Actions.BannerModule.GetByID.Do(ctx, req.GetBannerId())
		if err != nil {
			return err
		}

		resp = &desc.GetBannerByIDResponse{
			Banner: i.newBannerResponse(*banner),
		}

		return nil
	})
}
