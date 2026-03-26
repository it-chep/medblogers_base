package dto

import "time"

type OfferFilter struct {
	CooperationTypeIDs []int64
}

type NamedItem struct {
	ID   int64
	Name string
}

type SocialNetwork struct {
	ID   int64
	Name string
	Slug string
}

type Offer struct {
	Photo            string
	Title            string
	BrandDescription string
	CooperationType  *NamedItem
	Description      string
	SocialNetworks   []SocialNetwork
	BusinessCategory *NamedItem
	CreatedAt        *time.Time
}

type Response struct {
	Offers []Offer
}
