package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
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

func validateRequest(reqDTO create_doctor.CreateDoctorRequest) []error {
	validate := validator.New()

	err := validate.Struct(reqDTO)
	if err == nil {
		return nil
	}

	var validationErrors []error
	for _, err := range err.(validator.ValidationErrors) {
		var errMsg string
		switch err.Tag() {
		case "required":
			errMsg = fmt.Sprintf("Обязательное поле", err.Field())
		case "email":
			errMsg = fmt.Sprintf("Невалидный email", err.Field())
		case "max":
			errMsg = fmt.Sprintf("Текст нужно сократить", err.Field(), err.Param())
		default:
			errMsg = fmt.Sprintf("Неправильное значение", err.Field())
		}
		validationErrors = append(validationErrors, errors.New(errMsg))
	}

	return validationErrors
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
