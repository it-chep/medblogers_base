package dto

type DoctorItem struct {
	ID   int64
	Name string
	Slug string

	CityName       string
	SpecialityName string

	S3Image    string
	IsKFDoctor bool
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

type SearchDTO struct {
	Doctors      []DoctorItem
	Cities       []CityItem
	Specialities []SpecialityItem
}
