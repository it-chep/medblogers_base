package v1

import (
	"context"

	"github.com/google/uuid"

	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"
)

func (i *Implementation) GetPromoOfferSeoData(ctx context.Context, req *desc.GetPromoOfferSeoDataRequest) (*desc.GetPromoOfferSeoDataResponse, error) {
	offerID, err := uuid.Parse(req.GetOfferId())
	if err != nil {
		return nil, err
	}

	response, err := i.promoOffers.Actions.OfferSEO.Do(ctx, offerID)
	if err != nil {
		return nil, err
	}

	return &desc.GetPromoOfferSeoDataResponse{
		Title:       response.Title,
		Description: response.Description,
		Image:       response.ImageURL,
	}, nil
}
