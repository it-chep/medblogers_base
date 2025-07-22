package v1

import (
	"encoding/json"
	"medblogers_base/internal/app/api/doctors/v1/dto/search_doctor"
	"medblogers_base/internal/modules/doctors/action/search_doctor/dto"
	"net/http"
	"strings"

	"github.com/samber/lo"
)

// Search - /api/v1/doctors/search [GET]
func (s *Service) Search(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("query"))
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Получаем данные из модуля
	searchResultDomain, err := s.doctors.Actions.SearchDoctor.Do(r.Context(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Преобразуем в структуру с JSON тегами
	searchDTO := s.newSearchResponse(searchResultDomain)

	resp, err := json.Marshal(searchDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *Service) newSearchResponse(searchResultDomain dto.SearchDTO) search_doctor.SearchDTO {
	return search_doctor.SearchDTO{
		Doctors: lo.Map(searchResultDomain.Doctors, func(item dto.DoctorItem, _ int) search_doctor.DoctorItem {
			return search_doctor.DoctorItem{
				Name:           item.Name,
				Slug:           item.Slug,
				CityName:       item.CityName,
				SpecialityName: item.SpecialityName,
				S3Image:        item.S3Image,
			}
		}),
		Cities: lo.Map(searchResultDomain.Cities, func(cityItem dto.CityItem, _ int) search_doctor.CityItem {
			return search_doctor.CityItem{
				ID:           cityItem.ID,
				Name:         cityItem.Name,
				DoctorsCount: cityItem.DoctorsCount,
			}
		}),
		Specialities: lo.Map(searchResultDomain.Specialities, func(specialityItem dto.SpecialityItem, _ int) search_doctor.SpecialityItem {
			return search_doctor.SpecialityItem{
				ID:           specialityItem.ID,
				Name:         specialityItem.Name,
				DoctorsCount: specialityItem.DoctorsCount,
			}
		}),
	}
}
