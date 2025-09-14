package v1

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"medblogers_base/internal/app/api/freelancers/v1/validate/create_freelancer"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
	"reflect"
)

func (i *Implementation) CreateFreelancer(ctx context.Context, request *desc.CreateFreelancersRequest) (*desc.CreateFreelancersResponse, error) {
	requestErrors := validateRequest(request)
	if len(requestErrors) > 0 {
		return &desc.CreateFreelancersResponse{
			Errors: lo.Map(requestErrors, func(item create_freelancer.ValidationError, index int) *desc.CreateFreelancersResponse_ValidationError {
				return &desc.CreateFreelancersResponse_ValidationError{
					Text:  item.Text,
					Field: item.Field,
				}
			}),
		}, nil
	}

	createDTO := i.requestToCreateDoctorDTO(request)

	domainValidationErrors, err := i.freelancers.Actions.CreateFreelancer.Do(ctx, createDTO)
	if err != nil {
		return nil, err
	}

	if len(domainValidationErrors) > 0 {
		return i.configureResponse(domainValidationErrors), nil
	}

	return nil, nil
}

func (i *Implementation) requestToCreateDoctorDTO(req *desc.CreateFreelancersRequest) dto.CreateRequest {
	return dto.CreateRequest{
		Email:                 req.Email,
		LastName:              req.LastName,
		FirstName:             req.FirstName,
		MiddleName:            req.MiddleName,
		AdditionalCities:      req.AdditionalCities,
		AdditionalSpecialties: req.AdditionalSpecialties,
		MainCityID:            req.CityId,
		MainSpecialityID:      req.SpecialityId,
	}
}

func (i *Implementation) configureResponse(errors []dto.ValidationError) *desc.CreateFreelancersResponse {
	return &desc.CreateFreelancersResponse{
		Errors: lo.Map(errors, func(item dto.ValidationError, _ int) *desc.CreateFreelancersResponse_ValidationError {
			return &desc.CreateFreelancersResponse_ValidationError{
				Field: item.Field,
				Text:  item.Text,
			}
		}),
	}
}

func validateRequest(req *desc.CreateFreelancersRequest) []create_freelancer.ValidationError {
	reqDTO := create_freelancer.CreateFreelancerRequest{
		Email:            req.Email,
		LastName:         req.LastName,
		FirstName:        req.FirstName,
		MiddleName:       req.MiddleName,
		TelegramUsername: req.TelegramUsername,
		AgreePolicy:      req.AgreePolicy,
		CityID:           req.CityId,
		SpecialityID:     req.SpecialityId,
	}

	validate := validator.New()

	err := validate.Struct(reqDTO)
	if err == nil {
		return nil
	}

	var validationErrors []create_freelancer.ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		var validationError create_freelancer.ValidationError

		field := err.Field()
		if f, ok := reflect.TypeOf(create_freelancer.CreateFreelancerRequest{}).FieldByName(err.Field()); ok {
			if jsonTag := f.Tag.Get("json"); jsonTag != "" {
				field = jsonTag
			}
		}

		switch err.Tag() {
		case "required":
			validationError = create_freelancer.ValidationError{
				Field: field,
				Text:  "Обязательное поле",
			}
		case "email":
			validationError = create_freelancer.ValidationError{
				Field: field,
				Text:  "Невалидный email",
			}
		case "max":
			validationError = create_freelancer.ValidationError{
				Field: field,
				Text:  "Текст нужно сократить",
			}
		default:
			validationError = create_freelancer.ValidationError{
				Field: field,
				Text:  "Неправильное значение",
			}
		}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}
