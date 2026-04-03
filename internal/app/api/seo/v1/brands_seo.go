package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"
)

func (i *Implementation) GetBrandSeoData(ctx context.Context, req *desc.GetBrandSeoDataRequest) (*desc.GetBrandSeoDataResponse, error) {
	response, err := i.promoOffers.Actions.BrandSEO.Do(ctx, req.GetBrandSlug())
	if err != nil {
		return nil, err
	}

	return &desc.GetBrandSeoDataResponse{
		Title:       response.Title,
		Description: response.Description,
		Image:       response.ImageURL,
	}, nil
}
