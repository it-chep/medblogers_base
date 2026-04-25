package v1

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/offer/v1"
)

func (i *Implementation) DeactivateOffer(ctx context.Context, req *desc.DeactivateOfferRequest) (resp *desc.DeactivateOfferResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/offer/{id}/deactivate", func(ctx context.Context) error {
		offerID, err := uuid.Parse(req.GetOfferId())
		if err != nil {
			return err
		}

		if err := i.admin.Actions.PromoOffers.OfferAgg.DeactivateOffer.Do(ctx, offerID); err != nil {
			return err
		}

		resp = &desc.DeactivateOfferResponse{}
		return nil
	})
}
