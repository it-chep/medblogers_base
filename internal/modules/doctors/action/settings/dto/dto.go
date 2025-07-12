package dto

type FilterItem struct {
	Name string
	Slug string
}

type CityItem struct {
	ID           int64
	Name         string
	DoctorsCount int64
}

type SpecialityItem struct {
	ID           int64
	Name         string
	DoctorsCount int64
}

type Settings struct {
	DoctorsCount     int64
	SubscribersCount string
	FilterInfo       []FilterItem
	Cities           []CityItem
	Specialities     []SpecialityItem
	NewDoctorBanner  bool
}
