package dto

type NetworkItem struct {
	ID   int64
	Name string
}

type Freelancer struct {
	ID   int64
	Name string
	Slug string

	Speciality string // Строка из основной и дополнительных специальностей
	City       string // Строка из основного и дополнительных городов

	Image string

	MainCityID              int64
	MainSpecialityID        int64
	PriceCategory           int64
	HasExperienceWithDoctor bool

	Networks []NetworkItem
}

type Response struct {
	Freelancers []Freelancer
}
