package v1

import (
	"encoding/json"
	"github.com/it-chep/medblogers_base/internal/app/api/doctors/v1/dto"
	indto "github.com/it-chep/medblogers_base/internal/modules/doctors/action/settings/dto"
	"github.com/samber/lo"
	"net/http"
)

// Settings - /api/v1/settings [GET]
func (s *Service) Settings(w http.ResponseWriter, r *http.Request) {

	// Получаем данные из модуля
	settings, err := s.doctors.Actions.Settings.Do(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Преобразуем в структуру с JSON тегами
	settingsDTO := s.newSettingsResponse(settings)

	resp, err := json.Marshal(settingsDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *Service) newSettingsResponse(settings indto.Settings) dto.SettingsDTO {
	return dto.SettingsDTO{
		DoctorsCount:     settings.DoctorsCount,
		SubscribersCount: settings.SubscribersCount,
		NewDoctorBanner:  settings.NewDoctorBanner,
		FilterInfo: lo.Map(settings.FilterInfo, func(filterItem indto.FilterItem, _ int) dto.FilterItem {
			return dto.FilterItem{
				Name: filterItem.Name,
				Slug: filterItem.Slug,
			}
		}),
		Cities: lo.Map(settings.Cities, func(cityItem indto.CityItem, _ int) dto.CityItem {
			return dto.CityItem{
				ID:           cityItem.ID,
				Name:         cityItem.Name,
				DoctorsCount: cityItem.DoctorsCount,
			}
		}),
		Specialities: lo.Map(settings.Specialities, func(specialityItem indto.SpecialityItem, _ int) dto.SpecialityItem {
			return dto.SpecialityItem{
				ID:           specialityItem.ID,
				Name:         specialityItem.Name,
				DoctorsCount: specialityItem.DoctorsCount,
			}
		}),
	}
}
