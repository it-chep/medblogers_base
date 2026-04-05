package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

// GetBrandByOffer .
func (i *Implementation) GetBrandByOffer(ctx context.Context, req *desc.GetBrandByOfferRequest) (*desc.GetBrandByOfferResponse, error) {
	item, err := i.promoOffers.Actions.BrandByOffer.Do(ctx, req.GetOfferId())
	if err != nil {
		return nil, err
	}

	brand := &desc.BrandItem{
		Id:    item.ID,
		Title: item.Title,
		Slug:  item.Slug,
		Photo: item.Photo,
		BusinessCategory: &desc.NamedItem{
			Id:   item.BusinessCategory.ID,
			Name: item.BusinessCategory.Name,
		},
		SiteLink:       item.Website,
		Description:    item.Description,
		SocialNetworks: make([]*desc.BrandSocialNetworkItem, 0, len(item.SocialNetworks)),
	}

	for _, social := range item.SocialNetworks {
		brand.SocialNetworks = append(brand.SocialNetworks, &desc.BrandSocialNetworkItem{
			Id:   social.ID,
			Name: social.Name,
			Slug: social.Slug,
			Link: social.Link,
		})
	}

	return &desc.GetBrandByOfferResponse{Brand: brand}, nil
}
