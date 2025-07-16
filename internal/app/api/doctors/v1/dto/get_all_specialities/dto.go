package get_all_specialities

type SpecialityItem struct {
	ID   int64  `json:"speciality_id"`
	Name string `json:"speciality_name"`
}

type SpecialitiesResponse struct {
	Specialities []SpecialityItem `json:"specialities"`
}
