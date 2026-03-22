package v1

import (
	"context"

	"github.com/samber/lo"

	"medblogers_base/internal/app/interceptor"
	adminDto "medblogers_base/internal/modules/admin/entities/banner/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) GetBanners(ctx context.Context, _ *desc.GetBannersRequest) (resp *desc.GetBannersResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/banners", func(ctx context.Context) error {
		banners, err := i.admin.Actions.BannerModule.Get.Do(ctx)
		if err != nil {
			return err
		}

		resp = &desc.GetBannersResponse{
			Banners: lo.Map(banners, func(item adminDto.Banner, _ int) *desc.Banner {
				return i.newBannerResponse(item)
			}),
		}

		return nil
	})
}
