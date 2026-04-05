package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/dictionary/v1"
)

func (i *Implementation) CreateContentFormat(ctx context.Context, req *desc.CreateContentFormatRequest) (resp *desc.CreateContentFormatResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/content_formats/create", func(ctx context.Context) error {
		id, err := i.admin.Actions.PromoOffers.DictionaryAgg.CreateContentFormat.Do(ctx, req.GetName())
		if err != nil {
			return err
		}

		resp = newCreateContentFormatResponse(id)
		return nil
	})
}
