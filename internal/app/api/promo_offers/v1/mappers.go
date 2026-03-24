package v1

import (
	"time"

	brandDetailDTO "medblogers_base/internal/modules/promo_offers/action/brand_detail/dto"
	filterBrandsDTO "medblogers_base/internal/modules/promo_offers/action/filter_brands/dto"
	filterOffersDTO "medblogers_base/internal/modules/promo_offers/action/filter_offers/dto"
	filterSettingsDTO "medblogers_base/internal/modules/promo_offers/action/filter_settings/dto"
	offerDetailDTO "medblogers_base/internal/modules/promo_offers/action/offer_detail/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

func newFilterOffersResponse(resp filterOffersDTO.Response) *desc.FilterOffersResponse {
	result := &desc.FilterOffersResponse{
		Offers: make([]*desc.OfferItem, 0, len(resp.Offers)),
	}

	for _, item := range resp.Offers {
		result.Offers = append(result.Offers, newOfferItemFromFilter(item))
	}

	return result
}

func newFilterBrandsResponse(resp filterBrandsDTO.Response) *desc.FilterBrandsResponse {
	result := &desc.FilterBrandsResponse{
		Brands: make([]*desc.BrandItem, 0, len(resp.Brands)),
	}

	for _, item := range resp.Brands {
		result.Brands = append(result.Brands, newBrandItemFromFilter(item))
	}

	return result
}

func newGetBrandCardResponse(item *brandDetailDTO.Brand) *desc.GetBrandCardResponse {
	return &desc.GetBrandCardResponse{
		Brand: newBrandItemFromDetail(item),
	}
}

func newGetOfferCardResponse(item *offerDetailDTO.Offer) *desc.GetOfferCardResponse {
	return &desc.GetOfferCardResponse{
		Offer: newOfferItemFromDetail(item),
	}
}

func newGetFilterSettingsResponse(item *filterSettingsDTO.Response) *desc.GetFilterSettingsResponse {
	if item == nil {
		return nil
	}

	resp := &desc.GetFilterSettingsResponse{
		All:              item.All,
		CooperationTypes: make([]*desc.FilterCountItem, 0, len(item.CooperationTypes)),
	}

	for _, cooperationType := range item.CooperationTypes {
		resp.CooperationTypes = append(resp.CooperationTypes, &desc.FilterCountItem{
			Id:          cooperationType.ID,
			Name:        cooperationType.Name,
			OffersCount: cooperationType.OffersCount,
		})
	}

	return resp
}

func newOfferItemFromFilter(item filterOffersDTO.Offer) *desc.OfferItem {
	resp := &desc.OfferItem{
		Id:                   item.ID,
		Title:                item.Title,
		Description:          item.Description,
		Price:                item.Price,
		PublicationDate:      formatDate(item.PublicationDate),
		AdMarkingResponsible: item.AdMarkingResponsible,
		ResponsesCapacity:    item.ResponsesCapacity,
		CooperationType:      newNamedItemFromFilter(item.CooperationType),
		Topic:                newNamedItemFromFilter(item.Topic),
		ContentFormat:        newNamedItemFromFilter(item.ContentFormat),
		Brand:                newBrandPreviewFromFilter(item.Brand),
		SocialNetworks:       make([]*desc.SocialNetworkItem, 0, len(item.SocialNetworks)),
	}

	for _, social := range item.SocialNetworks {
		resp.SocialNetworks = append(resp.SocialNetworks, &desc.SocialNetworkItem{
			Id:   social.ID,
			Name: social.Name,
			Slug: social.Slug,
		})
	}

	return resp
}

func newOfferItemFromDetail(item *offerDetailDTO.Offer) *desc.OfferItem {
	resp := &desc.OfferItem{
		Id:                   item.ID,
		Title:                item.Title,
		Description:          item.Description,
		Price:                item.Price,
		PublicationDate:      formatDate(item.PublicationDate),
		AdMarkingResponsible: item.AdMarkingResponsible,
		ResponsesCapacity:    item.ResponsesCapacity,
		CooperationType:      newNamedItemFromDetail(item.CooperationType),
		Topic:                newNamedItemFromDetail(item.Topic),
		ContentFormat:        newNamedItemFromDetail(item.ContentFormat),
		Brand:                newBrandPreviewFromDetail(item.Brand),
		SocialNetworks:       make([]*desc.SocialNetworkItem, 0, len(item.SocialNetworks)),
	}

	for _, social := range item.SocialNetworks {
		resp.SocialNetworks = append(resp.SocialNetworks, &desc.SocialNetworkItem{
			Id:   social.ID,
			Name: social.Name,
			Slug: social.Slug,
		})
	}

	return resp
}

func newBrandItemFromFilter(item filterBrandsDTO.Brand) *desc.BrandItem {
	resp := &desc.BrandItem{
		Id:             item.ID,
		Title:          item.Title,
		Slug:           item.Slug,
		Photo:          item.Photo,
		Topic:          newTopicFromFilter(item.Topic),
		Website:        item.Website,
		Description:    item.Description,
		SocialNetworks: make([]*desc.BrandSocialNetworkItem, 0, len(item.SocialNetworks)),
	}

	for _, social := range item.SocialNetworks {
		resp.SocialNetworks = append(resp.SocialNetworks, &desc.BrandSocialNetworkItem{
			Id:   social.ID,
			Name: social.Name,
			Slug: social.Slug,
			Link: social.Link,
		})
	}

	return resp
}

func newBrandItemFromDetail(item *brandDetailDTO.Brand) *desc.BrandItem {
	if item == nil {
		return nil
	}

	resp := &desc.BrandItem{
		Id:             item.ID,
		Title:          item.Title,
		Slug:           item.Slug,
		Photo:          item.Photo,
		Topic:          newTopicFromDetail(item.Topic),
		Website:        item.Website,
		Description:    item.Description,
		SocialNetworks: make([]*desc.BrandSocialNetworkItem, 0, len(item.SocialNetworks)),
	}

	for _, social := range item.SocialNetworks {
		resp.SocialNetworks = append(resp.SocialNetworks, &desc.BrandSocialNetworkItem{
			Id:   social.ID,
			Name: social.Name,
			Slug: social.Slug,
			Link: social.Link,
		})
	}

	return resp
}

func newNamedItemFromFilter(item *filterOffersDTO.NamedItem) *desc.NamedItem {
	if item == nil {
		return nil
	}

	return &desc.NamedItem{
		Id:   item.ID,
		Name: item.Name,
	}
}

func newNamedItemFromDetail(item *offerDetailDTO.NamedItem) *desc.NamedItem {
	if item == nil {
		return nil
	}

	return &desc.NamedItem{
		Id:   item.ID,
		Name: item.Name,
	}
}

func newTopicFromFilter(item *filterBrandsDTO.Topic) *desc.NamedItem {
	if item == nil {
		return nil
	}

	return &desc.NamedItem{
		Id:   item.ID,
		Name: item.Name,
	}
}

func newTopicFromDetail(item *brandDetailDTO.Topic) *desc.NamedItem {
	if item == nil {
		return nil
	}

	return &desc.NamedItem{
		Id:   item.ID,
		Name: item.Name,
	}
}

func newBrandPreviewFromFilter(item *filterOffersDTO.BrandPreview) *desc.BrandPreview {
	if item == nil {
		return nil
	}

	return &desc.BrandPreview{
		Id:    item.ID,
		Title: item.Title,
		Slug:  item.Slug,
		Photo: item.Photo,
	}
}

func newBrandPreviewFromDetail(item *offerDetailDTO.BrandPreview) *desc.BrandPreview {
	if item == nil {
		return nil
	}

	return &desc.BrandPreview{
		Id:    item.ID,
		Title: item.Title,
		Slug:  item.Slug,
		Photo: item.Photo,
	}
}

func formatDate(value *time.Time) string {
	if value == nil {
		return ""
	}

	return value.Format(time.DateOnly)
}
