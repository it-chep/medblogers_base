package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/offer/v1"
)

func (i *Implementation) CreateOffer(ctx context.Context, req *desc.CreateOfferRequest) (resp *desc.CreateOfferResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/offers/create", func(ctx context.Context) error {
		offerID, err := i.admin.Actions.PromoOffers.OfferAgg.CreateOffer.Do(ctx, newCreateOfferDTO(req))
		if err != nil {
			return err
		}

		resp = &desc.CreateOfferResponse{OfferId: offerID.String()}
		return nil
	})
}
