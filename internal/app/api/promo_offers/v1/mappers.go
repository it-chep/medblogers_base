package v1

import (
	"github.com/samber/lo"
	brandDetailDTO "medblogers_base/internal/modules/promo_offers/action/brand_detail/dto"
	brandOffersDTO "medblogers_base/internal/modules/promo_offers/action/brand_offers/dto"
	filterBrandsDTO "medblogers_base/internal/modules/promo_offers/action/filter_brands/dto"
	filterOffersDTO "medblogers_base/internal/modules/promo_offers/action/filter_offers/dto"
	filterSettingsDTO "medblogers_base/internal/modules/promo_offers/action/filter_settings/dto"
	offerDetailDTO "medblogers_base/internal/modules/promo_offers/action/offer_detail/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
	"medblogers_base/internal/pkg/formatters"
)

func newFilterOffersResponse(resp filterOffersDTO.Response) *desc.FilterOffersResponse {
	result := &desc.FilterOffersResponse{
		Offers: make([]*desc.FilterOfferItem, 0, len(resp.Offers)),
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

func newGetBrandOffersResponse(resp brandOffersDTO.Response) *desc.GetBrandOffersResponse {
	result := &desc.GetBrandOffersResponse{
		Offers: make([]*desc.OfferItem, 0, len(resp.Offers)),
	}

	for _, item := range resp.Offers {
		result.Offers = append(result.Offers, newOfferItemFromBrandOffers(item))
	}

	return result
}

func newGetOfferCardResponse(item *offerDetailDTO.Offer) *desc.GetOfferCardResponse {
	return &desc.GetOfferCardResponse{
		Offer: newOfferCardItem(item),
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

func newOfferItemFromFilter(item filterOffersDTO.Offer) *desc.FilterOfferItem {
	resp := &desc.FilterOfferItem{
		Id:               item.ID,
		Photo:            item.Photo,
		Title:            item.Title,
		BrandDescription: item.BrandDescription,
		CooperationType:  newNamedItemFromFilter(item.CooperationType),
		Description:      item.Description,
		BusinessCategory: newNamedItemFromFilter(item.BusinessCategory),
		CreatedAt:        formatters.TimeRuFormat(item.CreatedAt),
		SocialNetworks:   make([]*desc.SocialNetworkItem, 0, len(item.SocialNetworks)),
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

func newOfferCardItem(item *offerDetailDTO.Offer) *desc.OfferCardItem {
	if item == nil {
		return nil
	}

	resp := &desc.OfferCardItem{
		Brand:           newOfferCardBrandItem(item.Brand),
		Description:     item.Description,
		CooperationType: newNamedItemFromDetail(item.CooperationType),
		Price:           item.Price,
		CreatedAt:       formatters.TimeRuFormat(item.CreatedAt),
		SocialNetworks:  make([]*desc.SocialNetworkItem, 0, len(item.SocialNetworks)),
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

func newOfferCardBrandItem(item *offerDetailDTO.Brand) *desc.OfferCardBrandItem {
	if item == nil {
		return nil
	}

	resp := &desc.OfferCardBrandItem{
		Photo:            item.Photo,
		Title:            item.Title,
		Description:      item.Description,
		About:            item.About,
		BusinessCategory: newNamedItemFromDetail(item.BusinessCategory),
		SocialNetworks:   make([]*desc.BrandSocialNetworkItem, 0, len(item.SocialNetworks)),
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

func newOfferItemFromBrandOffers(item brandOffersDTO.Offer) *desc.OfferItem {
	resp := &desc.OfferItem{
		Id:                   item.ID,
		Title:                item.Title,
		Description:          item.Description,
		Price:                item.Price,
		PublicationDate:      formatters.TimeRuFormat(lo.FromPtr(item.PublicationDate)),
		AdMarkingResponsible: item.AdMarkingResponsible,
		ResponsesCapacity:    item.ResponsesCapacity,
		CooperationType:      newNamedItemFromBrandOffers(item.CooperationType),
		BusinessCategory:     newNamedItemFromBrandOffers(item.BusinessCategory),
		ContentFormat:        newNamedItemFromBrandOffers(item.ContentFormat),
		Brand:                newBrandPreviewFromBrandOffers(item.Brand),
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
		Id:               item.ID,
		Title:            item.Title,
		Slug:             item.Slug,
		Photo:            item.Photo,
		BusinessCategory: newBusinessCategoryFromFilter(item.BusinessCategory),
		SiteLink:         item.Website,
		Description:      item.Description,
		SocialNetworks:   make([]*desc.BrandSocialNetworkItem, 0, len(item.SocialNetworks)),
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
		Id:               item.ID,
		Title:            item.Title,
		Slug:             item.Slug,
		Photo:            item.Photo,
		BusinessCategory: newBusinessCategoryFromDetail(item.BusinessCategory),
		SiteLink:         item.Website,
		Description:      item.Description,
		SocialNetworks:   make([]*desc.BrandSocialNetworkItem, 0, len(item.SocialNetworks)),
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

func newNamedItemFromBrandOffers(item *brandOffersDTO.NamedItem) *desc.NamedItem {
	if item == nil {
		return nil
	}

	return &desc.NamedItem{
		Id:   item.ID,
		Name: item.Name,
	}
}

func newBusinessCategoryFromFilter(item *filterBrandsDTO.BusinessCategory) *desc.NamedItem {
	if item == nil {
		return nil
	}

	return &desc.NamedItem{
		Id:   item.ID,
		Name: item.Name,
	}
}

func newBusinessCategoryFromDetail(item *brandDetailDTO.BusinessCategory) *desc.NamedItem {
	if item == nil {
		return nil
	}

	return &desc.NamedItem{
		Id:   item.ID,
		Name: item.Name,
	}
}

func newBrandPreviewFromBrandOffers(item *brandOffersDTO.BrandPreview) *desc.BrandPreview {
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
