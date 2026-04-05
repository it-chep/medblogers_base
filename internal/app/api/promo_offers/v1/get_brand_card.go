package v1

import (
	"context"

	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

func (i *Implementation) GetBrandCard(ctx context.Context, req *desc.GetBrandCardRequest) (*desc.GetBrandCardResponse, error) {
	resp, err := i.promoOffers.Actions.BrandDetail.Do(ctx, req.GetBrandSlug())
	if err != nil {
		return nil, err
	}

	return newGetBrandCardResponse(resp), nil
}
