package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) GetBrands(ctx context.Context, _ *desc.GetBrandsRequest) (resp *desc.GetBrandsResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brands", func(ctx context.Context) error {
		items, err := i.admin.Actions.PromoOffers.BrandAgg.GetBrands.Do(ctx)
		if err != nil {
			return err
		}

		resp = newGetBrandsResponse(items)
		return nil
	})
}
