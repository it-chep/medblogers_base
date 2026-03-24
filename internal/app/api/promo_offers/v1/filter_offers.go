package v1

import (
	"context"

	actionDTO "medblogers_base/internal/modules/promo_offers/action/filter_offers/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

func (i *Implementation) FilterOffers(ctx context.Context, req *desc.FilterOffersRequest) (*desc.FilterOffersResponse, error) {
	resp, err := i.promoOffers.Actions.FilterOffers.Do(ctx, actionDTO.OfferFilter{
		CooperationTypeIDs: req.GetCooperationTypeIds(),
	})
	if err != nil {
		return nil, err
	}

	return newFilterOffersResponse(resp), nil
}
