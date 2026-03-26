package v1

import (
	"time"

	createDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/create/dto"
	getDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get/dto"
	updateDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/update/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/brand/v1"
)

func newCreateBrandDTO(req *desc.CreateBrandRequest) createDTO.CreateRequest {
	return createDTO.CreateRequest{
		Title:              req.GetTitle(),
		Slug:               req.GetSlug(),
		BusinessCategoryID: req.GetBusinessCategoryId(),
		Website:            req.GetWebsite(),
		Description:        req.GetDescription(),
		SocialNetworks:     newCreateBrandSocialNetworks(req.GetSocialNetworks()),
	}
}

func newUpdateBrandDTO(req *desc.UpdateBrandRequest) updateDTO.UpdateRequest {
	return updateDTO.UpdateRequest{
		Title:              req.GetTitle(),
		Slug:               req.GetSlug(),
		BusinessCategoryID: req.GetBusinessCategoryId(),
		Website:            req.GetWebsite(),
		Description:        req.GetDescription(),
		SocialNetworks:     newUpdateBrandSocialNetworks(req.GetSocialNetworks()),
	}
}

func newGetBrandsResponse(items []getDTO.Brand) *desc.GetBrandsResponse {
	resp := &desc.GetBrandsResponse{
		Brands: make([]*desc.BrandListItem, 0, len(items)),
	}

	for _, item := range items {
		resp.Brands = append(resp.Brands, &desc.BrandListItem{
			Id:               item.ID,
			Title:            item.Title,
			Slug:             item.Slug,
			Photo:            item.Photo,
			BusinessCategory: newNamedItem(item.BusinessCategory),
			IsActive:         item.IsActive,
			CreatedAt:        formatDateTime(item.CreatedAt),
		})
	}

	return resp
}

func newGetBrandByIDResponse(item *getDTO.Brand) *desc.GetBrandByIDResponse {
	if item == nil {
		return &desc.GetBrandByIDResponse{}
	}

	resp := &desc.GetBrandByIDResponse{
		Brand: &desc.BrandItem{
			Id:               item.ID,
			Title:            item.Title,
			Slug:             item.Slug,
			Photo:            item.Photo,
			BusinessCategory: newNamedItem(item.BusinessCategory),
			Website:          item.Website,
			Description:      item.Description,
			SocialNetworks:   make([]*desc.BrandSocialNetworkItem, 0, len(item.SocialNetworks)),
			IsActive:         item.IsActive,
			CreatedAt:        formatDateTime(item.CreatedAt),
		},
	}

	for _, social := range item.SocialNetworks {
		resp.Brand.SocialNetworks = append(resp.Brand.SocialNetworks, &desc.BrandSocialNetworkItem{
			Id:   social.ID,
			Name: social.Name,
			Slug: social.Slug,
			Link: social.Link,
		})
	}

	return resp
}

func newCreateBrandSocialNetworks(items []*desc.BrandSocialNetworkInput) []createDTO.SocialNetworkInput {
	result := make([]createDTO.SocialNetworkInput, 0, len(items))
	for _, item := range items {
		result = append(result, createDTO.SocialNetworkInput{
			SocialNetworkID: item.GetSocialNetworkId(),
			Link:            item.GetLink(),
		})
	}

	return result
}

func newUpdateBrandSocialNetworks(items []*desc.BrandSocialNetworkInput) []updateDTO.SocialNetworkInput {
	result := make([]updateDTO.SocialNetworkInput, 0, len(items))
	for _, item := range items {
		result = append(result, updateDTO.SocialNetworkInput{
			SocialNetworkID: item.GetSocialNetworkId(),
			Link:            item.GetLink(),
		})
	}

	return result
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

func formatDateTime(value *time.Time) string {
	if value == nil {
		return ""
	}

	return value.Format(time.RFC3339)
}
