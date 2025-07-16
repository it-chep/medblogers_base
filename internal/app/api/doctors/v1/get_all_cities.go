package v1

import (
	"encoding/json"
	"github.com/samber/lo"
	"medblogers_base/internal/app/api/doctors/v1/dto/get_all_cities"
	"medblogers_base/internal/modules/doctors/domain/city"
	"net/http"
)

// AllCities - /api/v1/cities_list [GET]
func (s *Service) AllCities(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из модуля
	cities, err := s.doctors.Actions.AllCities.Do(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(s.getCitiesResponse(cities))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *Service) getCitiesResponse(citiesDomain []*city.City) get_all_cities.CitiesResponse {
	return get_all_cities.CitiesResponse{
		Cities: lo.Map(citiesDomain, func(item *city.City, _ int) get_all_cities.CityItem {
			return get_all_cities.CityItem{
				ID:   int64(item.ID()),
				Name: item.Name(),
			}
		}),
	}
}
