package v1

import (
	"context"

	actionDTO "medblogers_base/internal/modules/promo_offers/action/filter_brands/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

func (i *Implementation) FilterBrands(ctx context.Context, req *desc.FilterBrandsRequest) (*desc.FilterBrandsResponse, error) {
	resp, err := i.promoOffers.Actions.FilterBrands.Do(ctx, actionDTO.BrandFilter{
		TopicIDs:         req.GetTopicIds(),
		SocialNetworkIDs: req.GetSocialNetworkIds(),
	})
	if err != nil {
		return nil, err
	}

	return newFilterBrandsResponse(resp), nil
}
