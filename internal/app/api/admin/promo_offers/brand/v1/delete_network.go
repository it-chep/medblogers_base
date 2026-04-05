package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) DeleteNetwork(ctx context.Context, req *desc.DeleteNetworkRequest) (resp *desc.DeleteNetworkResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brand/{brand_id}/delete_network", func(ctx context.Context) error {
		if err := i.admin.Actions.PromoOffers.BrandAgg.DeleteNetwork.Do(ctx, req.GetBrandId(), req.GetSocialNetworkId()); err != nil {
			return err
		}

		resp = &desc.DeleteNetworkResponse{}
		return nil
	})
}
