package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) CreateBrand(ctx context.Context, req *desc.CreateBrandRequest) (resp *desc.CreateBrandResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brands/create", func(ctx context.Context) error {
		brandID, err := i.admin.Actions.PromoOffers.BrandAgg.CreateBrand.Do(ctx, newCreateBrandDTO(req))
		if err != nil {
			return err
		}

		resp = &desc.CreateBrandResponse{BrandId: brandID}
		return nil
	})
}
