package v1

import (
	"context"

	"github.com/google/uuid"

	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

func (i *Implementation) GetOfferCard(ctx context.Context, req *desc.GetOfferCardRequest) (*desc.GetOfferCardResponse, error) {
	offerID, err := uuid.Parse(req.GetOfferId())
	if err != nil {
		return nil, err
	}

	resp, err := i.promoOffers.Actions.OfferDetail.Do(ctx, offerID)
	if err != nil {
		return nil, err
	}

	return newGetOfferCardResponse(resp), nil
}
