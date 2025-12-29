package dto

type NetworkItem struct {
	ID   int64
	Name string
	Slug string
}

type Freelancer struct {
	ID   int64
	Name string
	Slug string

	Specialities []Speciality
	Cities       []City

	Image string

	MainCityID           int64
	MainSpecialityID     int64
	PriceCategory        int64
	AgencyRepresentative bool

	Networks []NetworkItem
}

type Speciality struct {
	ID   int64
	Name string
}

type City struct {
	ID   int64
	Name string
}
type Response struct {
	Freelancers []Freelancer
}
