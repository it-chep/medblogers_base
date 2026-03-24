package v1

import (
	"context"

	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

func (i *Implementation) GetBrandOffers(ctx context.Context, req *desc.GetBrandOffersRequest) (*desc.GetBrandOffersResponse, error) {
	resp, err := i.promoOffers.Actions.BrandOffers.Do(ctx, req.GetBrandSlug())
	if err != nil {
		return nil, err
	}

	return newGetBrandOffersResponse(resp), nil
}
