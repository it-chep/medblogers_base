package v1

import (
	"context"
	"encoding/base64"
	"log"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) SaveBannerDesktopImage(ctx context.Context, req *desc.SaveBannerDesktopImageRequest) (resp *desc.SaveBannerDesktopImageResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/banner/{id}/save_desktop_image", func(ctx context.Context) error {
		data, err := base64.StdEncoding.DecodeString(req.GetImageData())
		if err != nil {
			log.Fatal("Ошибка декодирования:", err)
			return err
		}

		imageURL, err := i.admin.Actions.BannerModule.SaveDesktopImage.Do(ctx, req.GetBannerId(), data)
		if err != nil {
			return err
		}

		resp = &desc.SaveBannerDesktopImageResponse{
			Image: imageURL,
		}

		return nil
	})
}
