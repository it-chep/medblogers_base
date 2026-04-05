package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/dictionary/v1"
)

func (i *Implementation) GetContentFormats(ctx context.Context, _ *desc.GetContentFormatsRequest) (resp *desc.GetContentFormatsResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/content_formats", func(ctx context.Context) error {
		items, err := i.admin.Actions.PromoOffers.DictionaryAgg.GetContentFormats.Do(ctx)
		if err != nil {
			return err
		}

		resp = newGetContentFormatsResponse(items)
		return nil
	})
}
