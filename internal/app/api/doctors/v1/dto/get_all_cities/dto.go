package get_all_cities

type CityItem struct {
	ID   int64  `json:"city_id"`
	Name string `json:"city_name"`
}

type CitiesResponse struct {
	Cities []CityItem `json:"cities"`
}
