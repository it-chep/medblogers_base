package dto

type BrandFilter struct {
	TopicIDs         []int64
	SocialNetworkIDs []int64
}

type Topic struct {
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
	ID             int64
	Title          string
	Slug           string
	Photo          string
	Topic          *Topic
	Website        string
	Description    string
	SocialNetworks []SocialNetwork
}

type Response struct {
	Brands []Brand
}
