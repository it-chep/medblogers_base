package rules

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
)

// RuleValidCityID проверяет валидность выбранного города
var RuleValidCityID = func(citiesIDs []int64) func(_ context.Context, req *dto.CreateRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateRequest) (bool, error) {
		if !lo.Contains(citiesIDs, req.MainCityID) {
			return false, dto.ValidationError{
				Text:  "Выбранного города не существует",
				Field: "cityId",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidAdditionalCitiesIDs проверяет валидность выбранных доп города
var RuleValidAdditionalCitiesIDs = func(citiesIDs []int64) func(_ context.Context, req *dto.CreateRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateRequest) (bool, error) {
		if len(req.AdditionalCities) == 0 {
			return true, dto.ValidationError{}
		}

		validCitiesMap := make(map[int64]struct{})
		for _, id := range citiesIDs {
			validCitiesMap[id] = struct{}{}
		}

		var invalidCities []int64
		for _, id := range req.AdditionalCities {
			if _, exists := validCitiesMap[id]; !exists {
				invalidCities = append(invalidCities, id)
			}
		}

		if len(invalidCities) > 0 {
			return false, dto.ValidationError{
				Text:  "Содержатся недопустимые дополнительные города",
				Field: "additionalCities",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidSpecialityID проверяет валидность выбранной специальности
var RuleValidSpecialityID = func(specialitiesIDs []int64) func(ctx context.Context, t *dto.CreateRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateRequest) (bool, error) {
		if !lo.Contains(specialitiesIDs, req.MainSpecialityID) {
			return false, dto.ValidationError{
				Text:  "Выбранной специальности не существует",
				Field: "specialityId",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidSpecialitiesIDs проверяет валидность выбранных доп специальностей
var RuleValidSpecialitiesIDs = func(specialitiesIDs []int64) func(ctx context.Context, t *dto.CreateRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateRequest) (bool, error) {
		if len(req.AdditionalCities) == 0 {
			return true, dto.ValidationError{}
		}

		validSpecialitiesMap := make(map[int64]struct{})
		for _, id := range specialitiesIDs {
			validSpecialitiesMap[id] = struct{}{}
		}

		var invalidSpecialities []int64
		for _, id := range req.AdditionalSpecialties {
			if _, exists := validSpecialitiesMap[id]; !exists {
				invalidSpecialities = append(invalidSpecialities, id)
			}
		}

		if len(invalidSpecialities) > 0 {
			return false, dto.ValidationError{
				Text:  "Содержатся недопустимые дополнительные специальности",
				Field: "additionalSpecialities",
			}
		}

		return true, dto.ValidationError{}
	}
}
