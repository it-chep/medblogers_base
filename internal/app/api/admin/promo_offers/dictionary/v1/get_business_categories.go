package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/dictionary/v1"
)

func (i *Implementation) GetBusinessCategories(ctx context.Context, _ *desc.GetBusinessCategoriesRequest) (resp *desc.GetBusinessCategoriesResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/business_categories", func(ctx context.Context) error {
		items, err := i.admin.Actions.PromoOffers.DictionaryAgg.GetTopics.Do(ctx)
		if err != nil {
			return err
		}

		resp = newGetBusinessCategoriesResponse(items)
		return nil
	})
}
