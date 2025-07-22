package v1

import (
	"encoding/json"
	"medblogers_base/internal/app/api/doctors/v1/dto/create_doctor"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"net/http"
)

// CreateDoctor /api/v1/doctors/create [POST]
func (s *Service) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var req create_doctor.CreateDoctorRequest

	// Декодируем JSON тело запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// todo валидация запроса по обязательным параметрам

	createDTO := s.requestToCreateDoctorDTO(req)
	err := s.doctors.Actions.CreateDoctor.Create(r.Context(), createDTO)
	if err != nil {
		return
	}

}

func (s *Service) requestToCreateDoctorDTO(reqDTO create_doctor.CreateDoctorRequest) dto.CreateDoctorRequest {
	return dto.CreateDoctorRequest{
		Email:                 reqDTO.Email,
		LastName:              reqDTO.LastName,
		FirstName:             reqDTO.FirstName,
		MiddleName:            reqDTO.MiddleName,
		BirthDate:             reqDTO.BirthDate,
		AdditionalCities:      reqDTO.AdditionalCities,
		AdditionalSpecialties: reqDTO.AdditionalSpecialties,
		InstagramUsername:     reqDTO.InstagramUsername,
		VKUsername:            reqDTO.VKUsername,
		TelegramUsername:      reqDTO.TelegramUsername,
		DzenUsername:          reqDTO.DzenUsername,
		YoutubeUsername:       reqDTO.YoutubeUsername,
		TelegramChannel:       reqDTO.TelegramChannel,
		CityID:                reqDTO.CityID,
		SpecialityID:          reqDTO.SpecialityID,
		MainBlogTheme:         reqDTO.MainBlogTheme,
		SiteLink:              reqDTO.SiteLink,
		AgreePolicy:           reqDTO.AgreePolicy,
	}
}
