package v1

import (
	"context"
	"medblogers_base/internal/app/api/doctors/v1/dto/create_doctor"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	"reflect"
	"strings"

	"github.com/samber/lo"

	"github.com/go-playground/validator/v10"
)

// CreateDoctor /api/v1/doctors/create [POST]
func (i *Implementation) CreateDoctor(ctx context.Context, req *desc.CreateDoctorRequest) (*desc.CreateDoctorResponse, error) {
	createDTO := i.requestToCreateDoctorDTO(req)

	domainValidationErrors, err := i.doctors.Actions.CreateDoctor.Create(ctx, createDTO)
	if err != nil {
		return nil, err
	}

	if len(domainValidationErrors) > 0 {
		return i.configureResponse(domainValidationErrors), nil
	}

	return nil, nil
}

func (i *Implementation) requestToCreateDoctorDTO(req *desc.CreateDoctorRequest) dto.CreateDoctorRequest {
	return dto.CreateDoctorRequest{
		Email:                 req.Email,
		LastName:              req.LastName,
		FirstName:             req.FirstName,
		MiddleName:            req.MiddleName,
		BirthDateString:       req.BirthDate,
		AdditionalCities:      req.AdditionalCities,
		AdditionalSpecialties: req.AdditionalSpecialties,
		InstagramUsername:     req.InstagramUsername,
		VKUsername:            req.VkUsername,
		TelegramUsername:      req.TelegramUsername,
		DzenUsername:          req.DzenUsername,
		YoutubeUsername:       req.YoutubeUsername,
		TelegramChannel:       req.TelegramChannel,
		CityID:                req.CityId,
		SpecialityID:          req.SpecialityId,
		MainBlogTheme:         req.MainBlogTheme,
		SiteLink:              req.SiteLink,
		AgreePolicy:           req.AgreePolicy,
	}
}

func (i *Implementation) configureResponse(errors []dto.ValidationError) *desc.CreateDoctorResponse {
	return &desc.CreateDoctorResponse{
		Errors: lo.Map(errors, func(item dto.ValidationError, _ int) *desc.CreateDoctorResponse_ValidationError {
			return &desc.CreateDoctorResponse_ValidationError{
				Field: item.Field,
				Text:  item.Text,
			}
		}),
	}
}

func validateRequest(reqDTO create_doctor.CreateDoctorRequest) []create_doctor.ValidationError {

	//requestErrors := validateRequest(req)
	//if len(requestErrors) > 0 {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusBadRequest)
	//	json.NewEncoder(w).Encode(create_doctor.Response{
	//		Errors: requestErrors,
	//	})
	//	return
	//}
	//
	validate := validator.New()

	err := validate.Struct(reqDTO)
	if err == nil {
		return nil
	}

	var validationErrors []create_doctor.ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		var validationError create_doctor.ValidationError

		field := err.Field()
		if f, ok := reflect.TypeOf(create_doctor.CreateDoctorRequest{}).FieldByName(err.Field()); ok {
			if jsonTag := f.Tag.Get("json"); jsonTag != "" {
				field = jsonTag
			}
		}

		switch err.Tag() {
		case "required":
			validationError = create_doctor.ValidationError{
				Field: strings.ToLower(field),
				Text:  "Обязательное поле",
			}
		case "email":
			validationError = create_doctor.ValidationError{
				Field: strings.ToLower(field),
				Text:  "Невалидный email",
			}
		case "max":
			validationError = create_doctor.ValidationError{
				Field: strings.ToLower(field),
				Text:  "Текст нужно сократить",
			}
		default:
			validationError = create_doctor.ValidationError{
				Field: strings.ToLower(field),
				Text:  "Неправильное значение",
			}
		}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}
