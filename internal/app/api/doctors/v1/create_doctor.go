package v1

import (
	"encoding/json"
	"github.com/samber/lo"
	"medblogers_base/internal/app/api/doctors/v1/dto/create_doctor"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// CreateDoctor /api/v1/doctors/create [POST]
func (s *Service) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var req create_doctor.CreateDoctorRequest

	// Декодируем JSON тело запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestErrors := validateRequest(req)
	if len(requestErrors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(create_doctor.Response{
			Errors: requestErrors,
		})
		return
	}

	createDTO := s.requestToCreateDoctorDTO(req)
	domainValidationErrors, err := s.doctors.Actions.CreateDoctor.Create(r.Context(), createDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(domainValidationErrors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		response := s.configureResponse(domainValidationErrors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(create_doctor.Response{})
}

func validateRequest(reqDTO create_doctor.CreateDoctorRequest) []create_doctor.ValidationError {
	validate := validator.New()

	err := validate.Struct(reqDTO)
	if err == nil {
		return nil
	}

	var validationErrors []create_doctor.ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		var validationError create_doctor.ValidationError
		switch err.Tag() {
		case "required":
			validationError = create_doctor.ValidationError{
				Field: err.Field(),
				Text:  "Обязательное поле",
			}
		case "email":
			validationError = create_doctor.ValidationError{
				Field: err.Field(),
				Text:  "Невалидный email",
			}
		case "max":
			validationError = create_doctor.ValidationError{
				Field: err.Field(),
				Text:  "Текст нужно сократить",
			}
		default:
			validationError = create_doctor.ValidationError{
				Field: err.Field(),
				Text:  "Неправильное значение",
			}
		}
		validationErrors = append(validationErrors, validationError)
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

func (s *Service) configureResponse(errors []dto.ValidationError) create_doctor.Response {
	return create_doctor.Response{
		Errors: lo.Map(errors, func(item dto.ValidationError, _ int) create_doctor.ValidationError {
			return create_doctor.ValidationError{
				Field: item.Field,
				Text:  item.Text,
			}
		}),
	}
}
