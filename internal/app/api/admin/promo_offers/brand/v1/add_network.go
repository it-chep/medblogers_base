package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func (i *Implementation) AddNetwork(ctx context.Context, req *desc.AddNetworkRequest) (resp *desc.AddNetworkResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/brand/{brand_id}/add_network", func(ctx context.Context) error {
		if err := i.admin.Actions.PromoOffers.BrandAgg.AddNetwork.Do(ctx, req.GetBrandId(), req.GetSocialNetworkId(), req.GetLink()); err != nil {
			return err
		}

		resp = &desc.AddNetworkResponse{}
		return nil
	})
}
