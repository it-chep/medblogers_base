package v1

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/offer/v1"
)

func (i *Implementation) UpdateOffer(ctx context.Context, req *desc.UpdateOfferRequest) (resp *desc.UpdateOfferResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/offer/{id}/update", func(ctx context.Context) error {
		offerID, err := uuid.Parse(req.GetOfferId())
		if err != nil {
			return err
		}

		if err := i.admin.Actions.PromoOffers.OfferAgg.UpdateOffer.Do(ctx, offerID, newUpdateOfferDTO(req)); err != nil {
			return err
		}

		resp = &desc.UpdateOfferResponse{}
		return nil
	})
}
