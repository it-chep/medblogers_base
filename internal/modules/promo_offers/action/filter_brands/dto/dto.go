package dto

type BrandFilter struct {
	BusinessCategoryIDs []int64
	SocialNetworkIDs    []int64
}

type BusinessCategory struct {
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
	BusinessCategory *BusinessCategory
	Website          string
	Description      string
	About            string
	SocialNetworks   []SocialNetwork
}

type Response struct {
	Brands []Brand
}
