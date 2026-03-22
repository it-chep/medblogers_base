package v1

import (
	"context"
	"encoding/base64"
	"log"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) SaveBannerMobileImage(ctx context.Context, req *desc.SaveBannerMobileImageRequest) (resp *desc.SaveBannerMobileImageResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/banner/{id}/save_mobile_image", func(ctx context.Context) error {
		data, err := base64.StdEncoding.DecodeString(req.GetImageData())
		if err != nil {
			log.Fatal("Ошибка декодирования:", err)
			return err
		}

		imageURL, err := i.admin.Actions.BannerModule.SaveMobileImage.Do(ctx, req.GetBannerId(), data)
		if err != nil {
			return err
		}

		resp = &desc.SaveBannerMobileImageResponse{
			Image: imageURL,
		}

		return nil
	})
}
