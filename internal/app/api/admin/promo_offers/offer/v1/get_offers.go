package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/offer/v1"
)

func (i *Implementation) GetOffers(ctx context.Context, req *desc.GetOffersRequest) (resp *desc.GetOffersResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/offers", func(ctx context.Context) error {
		items, err := i.admin.Actions.PromoOffers.OfferAgg.GetOffers.Do(ctx, newGetOffersRequest(req))
		if err != nil {
			return err
		}

		resp = newGetOffersResponse(items)
		return nil
	})
}
