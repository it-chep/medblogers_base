package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) GetBrandByID(ctx context.Context, req *desc.GetBrandByIDRequest) (resp *desc.GetBrandByIDResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brand/{brand_id}", func(ctx context.Context) error {
		item, err := i.admin.Actions.PromoOffers.BrandAgg.GetBrandByID.Do(ctx, req.GetBrandId())
		if err != nil {
			return err
		}

		resp = newGetBrandByIDResponse(item)
		return nil
	})
}
