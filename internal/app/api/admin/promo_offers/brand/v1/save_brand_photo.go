package v1

import (
	"context"
	"encoding/base64"
	"log"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) SaveBrandPhoto(ctx context.Context, req *desc.SaveBrandPhotoRequest) (resp *desc.SaveBrandPhotoResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brand/{brand_id}/save_photo", func(ctx context.Context) error {
		data, err := base64.StdEncoding.DecodeString(req.GetImageData())
		if err != nil {
			log.Fatal("Ошибка декодирования:", err)
			return err
		}

		imageURL, err := i.admin.Actions.PromoOffers.BrandAgg.SaveBrandPhoto.Do(ctx, req.GetBrandId(), data)
		if err != nil {
			return err
		}

		resp = &desc.SaveBrandPhotoResponse{
			Image: imageURL,
		}

		return nil
	})
}
