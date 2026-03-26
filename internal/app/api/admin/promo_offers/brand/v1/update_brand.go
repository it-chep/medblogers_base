package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) UpdateBrand(ctx context.Context, req *desc.UpdateBrandRequest) (resp *desc.UpdateBrandResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brand/{brand_id}/update", func(ctx context.Context) error {
		if err := i.admin.Actions.PromoOffers.BrandAgg.UpdateBrand.Do(ctx, req.GetBrandId(), newUpdateBrandDTO(req)); err != nil {
			return err
		}

		resp = &desc.UpdateBrandResponse{}
		return nil
	})
}
