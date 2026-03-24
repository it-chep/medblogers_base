package dto

import "time"

type OfferFilter struct {
	CooperationTypeIDs []int64
}

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

type Offer struct {
	ID                   string
	Title                string
	Description          string
	Price                int64
	PublicationDate      *time.Time
	AdMarkingResponsible string
	ResponsesCapacity    int64
	CooperationType      *NamedItem
	Topic                *NamedItem
	ContentFormat        *NamedItem
	Brand                *BrandPreview
	SocialNetworks       []SocialNetwork
}

type Response struct {
	Offers []Offer
}
