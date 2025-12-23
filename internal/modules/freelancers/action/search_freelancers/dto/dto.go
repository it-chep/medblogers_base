package dto

type SocialNetworkItem struct {
	ID   int64
	Name string
}

type FreelancerItem struct {
	ID   int64
	Name string
	Slug string

	CityName       string
	SpecialityName string

	S3Image string

	PriceCategory        int64
	AgencyRepresentative bool

	SocialNetworks []SocialNetworkItem
}

type CityItem struct {
	ID               int64
	Name             string
	FreelancersCount int64
}

type SpecialityItem struct {
	ID               int64
	Name             string
	FreelancersCount int64
}

type SearchDTO struct {
	Freelancers  []FreelancerItem
	Cities       []CityItem
	Specialities []SpecialityItem
}
