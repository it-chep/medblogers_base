package search_doctor

type DoctorItem struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`

	CityName       string `json:"city_name"`
	SpecialityName string `json:"speciality_name"`

	S3Image string `json:"image"`
}

type CityItem struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	DoctorsCount int64  `json:"doctors_count"`
}

type SpecialityItem struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	DoctorsCount int64  `json:"doctors_count"`
}

type SearchDTO struct {
	Doctors      []DoctorItem     `json:"doctors"`
	Cities       []CityItem       `json:"cities"`
	Specialities []SpecialityItem `json:"specialities"`
}
