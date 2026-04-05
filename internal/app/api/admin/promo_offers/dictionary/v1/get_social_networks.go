package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/dictionary/v1"
)

func (i *Implementation) GetSocialNetworks(ctx context.Context, _ *desc.GetSocialNetworksRequest) (resp *desc.GetSocialNetworksResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/social_networks", func(ctx context.Context) error {
		items, err := i.admin.Actions.PromoOffers.DictionaryAgg.GetSocialNetworks.Do(ctx)
		if err != nil {
			return err
		}

		resp = newGetSocialNetworksResponse(items)
		return nil
	})
}
