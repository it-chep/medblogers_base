package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) ActivateBrand(ctx context.Context, req *desc.ActivateBrandRequest) (resp *desc.ActivateBrandResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brand/{brand_id}/activate", func(ctx context.Context) error {
		if err := i.admin.Actions.PromoOffers.BrandAgg.ActivateBrand.Do(ctx, req.GetBrandId()); err != nil {
			return err
		}

		resp = &desc.ActivateBrandResponse{}
		return nil
	})
}
