package settings

// FilterItem - todo доделать
type FilterItem struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
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

type SettingsDTO struct {
	DoctorsCount     int64            `json:"doctors_count"`
	SubscribersCount string           `json:"subscribers_count"`
	FilterInfo       []FilterItem     `json:"filter_info"`
	Cities           []CityItem       `json:"cities"`
	Specialities     []SpecialityItem `json:"specialities"`
	NewDoctorBanner  bool             `json:"new_doctor_banner"`
}
