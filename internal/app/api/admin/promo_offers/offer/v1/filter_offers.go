package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/offer/v1"
)

func (i *Implementation) FilterOffers(ctx context.Context, req *desc.FilterOffersRequest) (resp *desc.FilterOffersResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/offers/filter", func(ctx context.Context) error {
		items, err := i.admin.Actions.PromoOffers.OfferAgg.FilterOffers.Do(ctx, newFilterOffersRequest(req))
		if err != nil {
			return err
		}

		resp = newFilterOffersResponse(items)
		return nil
	})
}
