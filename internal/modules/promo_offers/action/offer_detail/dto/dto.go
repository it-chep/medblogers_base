package dto

import "time"

type NamedItem struct {
	ID   int64
	Name string
}

type BrandPreview struct {
	ID    int64
	Title string
	Slug  string
	Photo string
}

type SocialNetwork struct {
	ID   int64
	Name string
	Slug string
}

type BrandSocialNetwork struct {
	ID   int64
	Name string
	Slug string
	Link string
}

type Brand struct {
	Photo            string
	Title            string
	Description      string
	About            string
	SocialNetworks   []BrandSocialNetwork
	BusinessCategory *NamedItem
}

type Offer struct {
	Brand           *Brand
	Description     string
	CooperationType *NamedItem
	Price           int64
	SocialNetworks  []SocialNetwork
	CreatedAt       time.Time
}
