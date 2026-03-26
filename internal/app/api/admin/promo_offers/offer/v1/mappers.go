package v1

import (
	"time"

	createDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/create/dto"
	filterDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/filter/dto"
	getDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get/dto"
	updateDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/update/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/offer/v1"
)

func newGetOffersRequest(req *desc.GetOffersRequest) getDTO.Request {
	return getDTO.Request{
		IsActive: req.IsActive,
		Page:     req.GetPage(),
		Limit:    req.GetLimit(),
	}
}

func newFilterOffersRequest(req *desc.FilterOffersRequest) filterDTO.Request {
	return filterDTO.Request{
		BrandIDs: req.GetBrandIds(),
		IsActive: req.IsActive,
		Page:     req.GetPage(),
		Limit:    req.GetLimit(),
	}
}

func newCreateOfferDTO(req *desc.CreateOfferRequest) createDTO.CreateRequest {
	return createDTO.CreateRequest{
		CooperationTypeID:    req.GetCooperationTypeId(),
		BusinessCategoryID:   req.GetBusinessCategoryId(),
		Title:                req.GetTitle(),
		Description:          req.GetDescription(),
		Price:                req.GetPrice(),
		ContentFormatID:      req.GetContentFormatId(),
		BrandID:              req.GetBrandId(),
		PublicationDate:      req.GetPublicationDate(),
		AdMarkingResponsible: req.GetAdMarkingResponsible(),
		ResponsesCapacity:    req.GetResponsesCapacity(),
		SocialNetworkIDs:     req.GetSocialNetworkIds(),
	}
}

func newUpdateOfferDTO(req *desc.UpdateOfferRequest) updateDTO.UpdateRequest {
	return updateDTO.UpdateRequest{
		CooperationTypeID:    req.GetCooperationTypeId(),
		BusinessCategoryID:   req.GetBusinessCategoryId(),
		Title:                req.GetTitle(),
		Description:          req.GetDescription(),
		Price:                req.GetPrice(),
		ContentFormatID:      req.GetContentFormatId(),
		BrandID:              req.GetBrandId(),
		PublicationDate:      req.GetPublicationDate(),
		AdMarkingResponsible: req.GetAdMarkingResponsible(),
		ResponsesCapacity:    req.GetResponsesCapacity(),
		SocialNetworkIDs:     req.GetSocialNetworkIds(),
	}
}

func newGetOffersResponse(items []getDTO.Offer) *desc.GetOffersResponse {
	resp := &desc.GetOffersResponse{
		Offers: make([]*desc.OfferListItem, 0, len(items)),
	}

	for _, item := range items {
		resp.Offers = append(resp.Offers, &desc.OfferListItem{
			Id:                item.ID,
			Title:             item.Title,
			Price:             item.Price,
			PublicationDate:   formatDateTime(item.PublicationDate),
			ResponsesCapacity: item.ResponsesCapacity,
			CooperationType:   newNamedItem(item.CooperationType),
			BusinessCategory:  newNamedItem(item.BusinessCategory),
			ContentFormat:     newNamedItem(item.ContentFormat),
			Brand:             newBrandPreview(item.Brand),
			IsActive:          item.IsActive,
			CreatedAt:         formatDateTime(item.CreatedAt),
		})
	}

	return resp
}

func newFilterOffersResponse(items []getDTO.Offer) *desc.FilterOffersResponse {
	resp := &desc.FilterOffersResponse{
		Offers: make([]*desc.OfferListItem, 0, len(items)),
	}

	for _, item := range newGetOffersResponse(items).GetOffers() {
		resp.Offers = append(resp.Offers, item)
	}

	return resp
}

func newGetOfferByIDResponse(item *getDTO.Offer) *desc.GetOfferByIDResponse {
	if item == nil {
		return &desc.GetOfferByIDResponse{}
	}

	resp := &desc.GetOfferByIDResponse{
		Offer: &desc.OfferItem{
			Id:                   item.ID,
			Title:                item.Title,
			Description:          item.Description,
			Price:                item.Price,
			PublicationDate:      formatDateTime(item.PublicationDate),
			AdMarkingResponsible: item.AdMarkingResponsible,
			ResponsesCapacity:    item.ResponsesCapacity,
			CooperationType:      newNamedItem(item.CooperationType),
			BusinessCategory:     newNamedItem(item.BusinessCategory),
			ContentFormat:        newNamedItem(item.ContentFormat),
			Brand:                newBrandPreview(item.Brand),
			SocialNetworks:       make([]*desc.SocialNetworkItem, 0, len(item.SocialNetworks)),
			IsActive:             item.IsActive,
			CreatedAt:            formatDateTime(item.CreatedAt),
		},
	}

	for _, social := range item.SocialNetworks {
		resp.Offer.SocialNetworks = append(resp.Offer.SocialNetworks, &desc.SocialNetworkItem{
			Id:   social.ID,
			Name: social.Name,
			Slug: social.Slug,
		})
	}

	return resp
}

func newNamedItem(item *getDTO.NamedItem) *desc.NamedItem {
	if item == nil {
		return nil
	}

	return &desc.NamedItem{
		Id:   item.ID,
		Name: item.Name,
	}
}

func newBrandPreview(item *getDTO.BrandPreview) *desc.BrandPreview {
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

func formatDateTime(value *time.Time) string {
	if value == nil {
		return ""
	}

	return value.Format(time.RFC3339)
}
