package dao

import (
	"github.com/google/uuid"

	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
)

type NamedDAO struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (d NamedDAO) ToDomain() *dictionary.NamedItem {
	return dictionary.NewNamedItem(
		dictionary.WithNamedItemID(d.ID),
		dictionary.WithNamedItemName(d.Name),
	)
}

type SocialNetworkDAO struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Slug string `db:"slug"`
}

func (d SocialNetworkDAO) ToDomain() *dictionary.SocialNetwork {
	return dictionary.NewSocialNetwork(
		dictionary.WithSocialNetworkID(d.ID),
		dictionary.WithSocialNetworkName(d.Name),
		dictionary.WithSocialNetworkSlug(d.Slug),
	)
}

type BrandSocialNetworkLinkDAO struct {
	BrandID         int64  `db:"brand_id"`
	SocialNetworkID int64  `db:"social_network_id"`
	Link            string `db:"link"`
}

type BrandSocialNetworkDAO struct {
	BrandID         int64
	SocialNetworkID int64
	Name            string
	Slug            string
	Link            string
}

func (d BrandSocialNetworkDAO) ToDomain() *dictionary.BrandSocialNetwork {
	return dictionary.NewBrandSocialNetwork(
		dictionary.WithBrandSocialNetworkID(d.SocialNetworkID),
		dictionary.WithBrandSocialNetworkName(d.Name),
		dictionary.WithBrandSocialNetworkSlug(d.Slug),
		dictionary.WithBrandSocialNetworkLink(d.Link),
	)
}

type OfferSocialNetworkLinkDAO struct {
	OfferID         uuid.UUID `db:"promo_offer_id"`
	SocialNetworkID int64     `db:"social_network_id"`
}

type OfferSocialNetworkDAO struct {
	OfferID         uuid.UUID
	SocialNetworkID int64
	Name            string
	Slug            string
}
