package v1

import (
	"encoding/json"
	"medblogers_base/internal/app/api/doctors/v1/dto/doctors_filter"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	indto "medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"net/http"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

// Filter - /api/v1/doctors/filter [GET]
func (s *Service) Filter(w http.ResponseWriter, r *http.Request) {
	filter := s.requestToFilterDTO(r)
	// Получаем данные из модуля
	filterResultDomain, err := s.doctors.Actions.DoctorsFilter.Do(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Преобразуем в структуру с JSON тегами
	filterDTO := s.newFilterResponse(filterResultDomain)

	resp, err := json.Marshal(filterDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *Service) requestToFilterDTO(r *http.Request) indto.Filter {
	maxSubscribers, _ := strconv.Atoi(r.URL.Query().Get("max_subscribers"))
	minSubscribers, _ := strconv.Atoi(r.URL.Query().Get("min_subscribers"))
	socialMedia := strings.Split(r.URL.Query().Get("social_media"), ",")

	cityParam := r.URL.Query().Get("city")
	var uniqueCities []int64
	if len(cityParam) != 0 {
		citiesStrings := strings.Split(r.URL.Query().Get("city"), ",")
		cities := lo.Map(citiesStrings, func(cityString string, _ int) int64 {
			cityID, err := strconv.Atoi(cityString)
			if err != nil {
				return 0
			}
			return int64(cityID)
		})
		uniqueCities = lo.Uniq(cities)
	}

	specialityParam := r.URL.Query().Get("speciality")
	var uniqueSpec []int64
	if len(specialityParam) != 0 {
		specialitiesStrings := strings.Split(r.URL.Query().Get("speciality"), ",")
		specialities := lo.Map(specialitiesStrings, func(specialityString string, _ int) int64 {
			specialityID, err := strconv.Atoi(specialityString)
			if err != nil {
				return 0
			}
			return int64(specialityID)
		})
		uniqueSpec = lo.Uniq(specialities)
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	return indto.Filter{
		MaxSubscribers: int64(maxSubscribers),
		MinSubscribers: int64(minSubscribers),
		Page:           int64(page),
		Cities:         uniqueCities,
		Specialities:   uniqueSpec,
		SocialMedia:    socialMedia,
	}
}
func (s *Service) newFilterResponse(filterDomain dto.Response) doctors_filter.Response {
	return doctors_filter.Response{
		Doctors: lo.Map(filterDomain.Doctors, func(item dto.Doctor, _ int) doctors_filter.DoctorItem {
			return doctors_filter.DoctorItem{
				ID:                item.ID,
				Name:              item.Name,
				Slug:              item.Slug,
				InstLink:          item.InstLink,
				InstSubsCount:     item.InstSubsCount,
				InstSubsCountText: item.InstSubsCountText,
				TgLink:            item.TgLink,
				TgSubsCount:       item.TgSubsCount,
				TgSubsCountText:   item.TgSubsCountText,
				Speciality:        item.Speciality,
				City:              item.City,
				Image:             item.Image,
			}
		}),
		CurrentPage: filterDomain.CurrentPage,
		Pages:       filterDomain.Pages,
	}
}
