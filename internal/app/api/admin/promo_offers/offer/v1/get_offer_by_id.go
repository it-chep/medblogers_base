package v1

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/offer/v1"
)

func (i *Implementation) GetOfferByID(ctx context.Context, req *desc.GetOfferByIDRequest) (resp *desc.GetOfferByIDResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/offer/{offer_id}", func(ctx context.Context) error {
		offerID, err := uuid.Parse(req.GetOfferId())
		if err != nil {
			return err
		}

		item, err := i.admin.Actions.PromoOffers.OfferAgg.GetOfferByID.Do(ctx, offerID)
		if err != nil {
			return err
		}

		resp = newGetOfferByIDResponse(item)
		return nil
	})
}
