package v1

import (
	"encoding/json"
	"medblogers_base/internal/app/api/doctors/v1/dto/doctor_detail"
	indto "medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"net/http"
	"strconv"

	"github.com/samber/lo"

	"github.com/go-chi/chi/v5"
)

// DoctorDetail - /api/v1/doctors/{doctor_id} [GET]
func (s *Service) DoctorDetail(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из URL
	idStr := chi.URLParam(r, "doctor_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}

	// Получаем данные из модуля
	doctorDomain, err := s.doctors.Actions.DoctorDetail.Do(r.Context(), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	doctorDTO := s.newDoctorDetailResponse(doctorDomain)
	resp, err := json.Marshal(doctorDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *Service) newDoctorDetailResponse(doctorDomain *indto.DoctorDTO) doctor_detail.DoctorDTO {
	return doctor_detail.DoctorDTO{
		Name: doctorDomain.Name,
		Slug: doctorDomain.Slug,

		InstURL:      doctorDomain.InstURL,
		VkURL:        doctorDomain.VkURL,
		DzenURL:      doctorDomain.DzenURL,
		YoutubeURL:   doctorDomain.YoutubeURL,
		TgURL:        doctorDomain.TgURL,
		TgChannelURL: doctorDomain.TgChannelURL,
		TiktokURL:    doctorDomain.TiktokURL,
		SiteLink:     doctorDomain.SiteLink,

		TgLastUpdatedDate: doctorDomain.TgLastUpdatedDate,
		TgSubsCountText:   doctorDomain.TgSubsCountText,
		TgSubsCount:       doctorDomain.TgSubsCount,

		InstSubsCount:       doctorDomain.InstSubsCount,
		InstSubsCountText:   doctorDomain.InstSubsCountText,
		InstLastUpdatedDate: doctorDomain.InstLastUpdatedDate,

		Cities: lo.Map(doctorDomain.Cities, func(item indto.CityItem, _ int) doctor_detail.CityItem {
			return doctor_detail.CityItem{
				ID:   item.ID,
				Name: item.Name,
			}
		}),

		Specialities: lo.Map(doctorDomain.Specialities, func(item indto.SpecialityItem, _ int) doctor_detail.SpecialityItem {
			return doctor_detail.SpecialityItem{
				ID:   item.ID,
				Name: item.Name,
			}
		}),

		MainBlogTheme:    doctorDomain.MainBlogTheme,
		MedicalDirection: doctorDomain.MedicalDirection,
		Image:            doctorDomain.Image,
	}
}
