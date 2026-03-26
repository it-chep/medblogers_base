package v1

import (
	domainDictionary "medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/dictionary/v1"
)

func newGetBusinessCategoriesResponse(items domainDictionary.NamedItems) *desc.GetBusinessCategoriesResponse {
	return &desc.GetBusinessCategoriesResponse{
		BusinessCategories: newNamedItems(items),
	}
}

func newCreateBusinessCategoryResponse(id int64) *desc.CreateBusinessCategoryResponse {
	return &desc.CreateBusinessCategoryResponse{Id: id}
}

func newGetCooperationTypesResponse(items domainDictionary.NamedItems) *desc.GetCooperationTypesResponse {
	return &desc.GetCooperationTypesResponse{
		CooperationTypes: newNamedItems(items),
	}
}

func newCreateCooperationTypeResponse(id int64) *desc.CreateCooperationTypeResponse {
	return &desc.CreateCooperationTypeResponse{Id: id}
}

func newGetContentFormatsResponse(items domainDictionary.NamedItems) *desc.GetContentFormatsResponse {
	return &desc.GetContentFormatsResponse{
		ContentFormats: newNamedItems(items),
	}
}

func newCreateContentFormatResponse(id int64) *desc.CreateContentFormatResponse {
	return &desc.CreateContentFormatResponse{Id: id}
}

func newGetSocialNetworksResponse(items domainDictionary.SocialNetworks) *desc.GetSocialNetworksResponse {
	return &desc.GetSocialNetworksResponse{
		SocialNetworks: newSocialNetworkItems(items),
	}
}

func newNamedItems(items domainDictionary.NamedItems) []*desc.NamedItem {
	result := make([]*desc.NamedItem, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}

		result = append(result, &desc.NamedItem{
			Id:   item.ID(),
			Name: item.Name(),
		})
	}

	return result
}

func newSocialNetworkItems(items domainDictionary.SocialNetworks) []*desc.SocialNetworkItem {
	result := make([]*desc.SocialNetworkItem, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}

		result = append(result, &desc.SocialNetworkItem{
			Id:   item.ID(),
			Name: item.Name(),
			Slug: item.Slug(),
		})
	}

	return result
}
