package v1

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/offer/v1"
)

func (i *Implementation) ActivateOffer(ctx context.Context, req *desc.ActivateOfferRequest) (resp *desc.ActivateOfferResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/offer/{offer_id}/activate", func(ctx context.Context) error {
		offerID, err := uuid.Parse(req.GetOfferId())
		if err != nil {
			return err
		}

		if err := i.admin.Actions.PromoOffers.OfferAgg.ActivateOffer.Do(ctx, offerID); err != nil {
			return err
		}

		resp = &desc.ActivateOfferResponse{}
		return nil
	})
}
