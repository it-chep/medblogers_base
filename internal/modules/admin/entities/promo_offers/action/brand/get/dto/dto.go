package dto

import (
	"time"

	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
)

type NamedItem struct {
	ID   int64
	Name string
}

type SocialNetwork struct {
	ID   int64
	Name string
	Slug string
	Link string
}

type Brand struct {
	ID               int64
	Title            string
	Slug             string
	Photo            string
	BusinessCategory *NamedItem
	Website          string
	Description      string
	About            string
	SocialNetworks   []SocialNetwork
	IsActive         bool
	CreatedAt        *time.Time
}

func NewBrand(item *brandDomain.Brand, businessCategoryName string, socials dictionary.BrandSocialNetworks, photo string) Brand {
	result := Brand{
		ID:             item.GetID(),
		Title:          item.GetTitle(),
		Slug:           item.GetSlug(),
		Photo:          photo,
		Website:        item.GetWebsite(),
		Description:    item.GetDescription(),
		About:          item.GetAbout(),
		SocialNetworks: make([]SocialNetwork, 0, len(socials)),
		IsActive:       item.GetIsActive(),
		CreatedAt:      item.GetCreatedAt(),
	}

	if item.GetBusinessCategoryID() > 0 && businessCategoryName != "" {
		result.BusinessCategory = &NamedItem{
			ID:   item.GetBusinessCategoryID(),
			Name: businessCategoryName,
		}
	}

	for _, social := range socials {
		result.SocialNetworks = append(result.SocialNetworks, SocialNetwork{
			ID:   social.SocialNetworkID(),
			Name: social.Name(),
			Slug: social.Slug(),
			Link: social.Link(),
		})
	}

	return result
}
