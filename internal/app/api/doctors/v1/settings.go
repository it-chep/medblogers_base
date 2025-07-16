package v1

import (
	"encoding/json"
	"medblogers_base/internal/app/api/doctors/v1/dto/settings"
	indto "medblogers_base/internal/modules/doctors/action/settings/dto"
	"net/http"

	"github.com/samber/lo"
)

// Settings - /api/v1/settings [GET]
func (s *Service) Settings(w http.ResponseWriter, r *http.Request) {

	// Получаем данные из модуля
	settingsDomain, err := s.doctors.Actions.Settings.Do(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Преобразуем в структуру с JSON тегами
	settingsDTO := s.newSettingsResponse(settingsDomain)

	resp, err := json.Marshal(settingsDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *Service) newSettingsResponse(settingsDomain *indto.Settings) settings.SettingsDTO {
	return settings.SettingsDTO{
		DoctorsCount:     settingsDomain.DoctorsCount,
		SubscribersCount: settingsDomain.SubscribersCount,
		NewDoctorBanner:  settingsDomain.NewDoctorBanner,
		FilterInfo: lo.Map(settingsDomain.FilterInfo, func(filterItem indto.FilterItem, _ int) settings.FilterItem {
			return settings.FilterItem{
				Name: filterItem.Name,
				Slug: filterItem.Slug,
			}
		}),
		Cities: lo.Map(settingsDomain.Cities, func(cityItem indto.CityItem, _ int) settings.CityItem {
			return settings.CityItem{
				ID:           cityItem.ID,
				Name:         cityItem.Name,
				DoctorsCount: cityItem.DoctorsCount,
			}
		}),
		Specialities: lo.Map(settingsDomain.Specialities, func(specialityItem indto.SpecialityItem, _ int) settings.SpecialityItem {
			return settings.SpecialityItem{
				ID:           specialityItem.ID,
				Name:         specialityItem.Name,
				DoctorsCount: specialityItem.DoctorsCount,
			}
		}),
	}
}
