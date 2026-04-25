package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) DeactivateBrand(ctx context.Context, req *desc.DeactivateBrandRequest) (resp *desc.DeactivateBrandResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brand/{id}/deactivate", func(ctx context.Context) error {
		if err := i.admin.Actions.PromoOffers.BrandAgg.DeactivateBrand.Do(ctx, req.GetBrandId()); err != nil {
			return err
		}

		resp = &desc.DeactivateBrandResponse{}
		return nil
	})
}
