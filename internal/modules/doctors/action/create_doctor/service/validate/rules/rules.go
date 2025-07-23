package rules

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/pkg/spec"
	"regexp"
	"strings"
	"time"

	"github.com/samber/lo"
)

// RuleAtLeastOneSocialMedia проверяет что у пользователя указана хотя бы 1 соцсеть
var RuleAtLeastOneSocialMedia = func() func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
		if req.InstagramUsername == "" &&
			req.VKUsername == "" &&
			req.TelegramUsername == "" &&
			req.TelegramChannel == "" &&
			req.DzenUsername == "" &&
			req.YoutubeUsername == "" {
			return false, spec.SpecError{
				Message: "обязательно нужно указать хотя бы 1 вашу соцсеть",
				Field:   "instagram_username",
			}
		}
		return true, spec.SpecError{}
	}
}

// RuleValidCityID проверяет валидность выбранного города
var RuleValidCityID = func(citiesIDs []int64) func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
		if !lo.Contains(citiesIDs, req.CityID) {
			return false, spec.SpecError{
				Message: "Выбранного города не существует",
				Field:   "city_id",
			}
		}

		return true, spec.SpecError{}
	}
}

// RuleValidAdditionalCitiesIDs проверяет валидность выбранных доп города
var RuleValidAdditionalCitiesIDs = func(citiesIDs []int64) func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
		if len(req.AdditionalCities) == 0 {
			return true, spec.SpecError{}
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
			return false, spec.SpecError{
				Message: "содержатся недопустимые ID дополнительных городов",
				Field:   "additional_cities",
			}
		}

		return true, spec.SpecError{}
	}
}

// RuleValidSpecialityID проверяет валидность выбранной специальности
var RuleValidSpecialityID = func(specialitiesIDs []int64) func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, spec.SpecError) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
		if !lo.Contains(specialitiesIDs, req.SpecialityID) {
			return false, spec.SpecError{
				Message: "Выбранной специальности не существует",
				Field:   "speciality_id",
			}
		}

		return true, spec.SpecError{}
	}
}

// RuleValidSpecialitiesIDs проверяет валидность выбранных доп специальностей
var RuleValidSpecialitiesIDs = func(specialitiesIDs []int64) func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, spec.SpecError) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
		if len(req.AdditionalCities) == 0 {
			return true, spec.SpecError{}
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
			return false, spec.SpecError{
				Message: "содержатся недопустимые ID дополнительных специальностей",
				Field:   "additional_specialities",
			}
		}

		return true, spec.SpecError{}
	}
}

// RuleValidSiteLink проверяет валидность сайта
var RuleValidSiteLink = func() func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, spec.SpecError) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
		if !strings.HasPrefix(req.SiteLink, "http") {
			return false, spec.SpecError{
				Message: "пожалуйста, укажите ссылку на сайт, ссылка должна содержать http",
				Field:   "site_link",
			}
		}

		return true, spec.SpecError{}
	}
}

// RuleValidBirthDate проверяет валидность дня рождения
var RuleValidBirthDate = func() func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, spec.SpecError) {
		birthDateStr := req.BirthDate

		if birthDateStr == "" {
			return false, spec.SpecError{
				Message: "дата рождения обязательна",
				Field:   "birth_date",
			}
		}

		dateRegex := regexp.MustCompile(`^\d{2}\.\d{2}\.\d{4}$`)
		if !dateRegex.MatchString(birthDateStr) {
			return false, spec.SpecError{
				Message: "неверный формат даты. ожидается ДД.ММ.ГГГГ",
				Field:   "birth_date",
			}
		}

		birthDate, err := time.Parse("02.01.2006", birthDateStr)
		if err != nil {
			return false, spec.SpecError{
				Message: "некорректная дата рождения",
				Field:   "birth_date",
			}
		}

		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		if birthDate.After(today) {
			return false, spec.SpecError{
				Message: "дата рождения не может быть в будущем",
				Field:   "birth_date",
			}
		}

		maxAgeDate := today.AddDate(-120, 0, 0)
		if birthDate.Before(maxAgeDate) {
			return false, spec.SpecError{
				Message: "дата рождения не должна быть старше 120 лет",
				Field:   "birth_date",
			}
		}

		return true, spec.SpecError{}
	}
}
