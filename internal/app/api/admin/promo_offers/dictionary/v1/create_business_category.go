package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/dictionary/v1"
)

func (i *Implementation) CreateBusinessCategory(ctx context.Context, req *desc.CreateBusinessCategoryRequest) (resp *desc.CreateBusinessCategoryResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/business_categories/create", func(ctx context.Context) error {
		id, err := i.admin.Actions.PromoOffers.DictionaryAgg.CreateTopic.Do(ctx, req.GetName())
		if err != nil {
			return err
		}

		resp = newCreateBusinessCategoryResponse(id)
		return nil
	})
}
