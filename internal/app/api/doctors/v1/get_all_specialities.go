package v1

import (
	"encoding/json"
	"medblogers_base/internal/app/api/doctors/v1/dto/get_all_specialities"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"net/http"

	"github.com/samber/lo"
)

// AllSpecialities - /api/v1/specialities_list [GET]
func (s *Service) AllSpecialities(w http.ResponseWriter, r *http.Request) {

	// Получаем данные из модуля
	specialities, err := s.doctors.Actions.AllSpecialities.Do(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(s.getSpecialitiesResponse(specialities))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *Service) getSpecialitiesResponse(specialitiesDomain []*speciality.Speciality) get_all_specialities.SpecialitiesResponse {
	return get_all_specialities.SpecialitiesResponse{
		Specialities: lo.Map(specialitiesDomain, func(item *speciality.Speciality, _ int) get_all_specialities.SpecialityItem {
			return get_all_specialities.SpecialityItem{
				ID:   int64(item.ID()),
				Name: item.Name(),
			}
		}),
	}
}
