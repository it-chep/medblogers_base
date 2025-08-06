package rules

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"regexp"
	"strings"
	"time"

	"github.com/samber/lo"
)

// RuleAtLeastOneSocialMedia проверяет что у пользователя указана хотя бы 1 соцсеть
var RuleAtLeastOneSocialMedia = func() func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if req.InstagramUsername == "" &&
			req.VKUsername == "" &&
			req.TelegramUsername == "" &&
			req.TelegramChannel == "" &&
			req.DzenUsername == "" &&
			req.YoutubeUsername == "" {
			return false, dto.ValidationError{
				Text:  "Обязательно нужно указать хотя бы 1 вашу соцсеть",
				Field: "instagram_username",
			}
		}
		return true, dto.ValidationError{}
	}
}

// RuleValidCityID проверяет валидность выбранного города
var RuleValidCityID = func(citiesIDs []int64) func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if !lo.Contains(citiesIDs, req.CityID) {
			return false, dto.ValidationError{
				Text:  "Выбранного города не существует",
				Field: "city_id",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidAdditionalCitiesIDs проверяет валидность выбранных доп города
var RuleValidAdditionalCitiesIDs = func(citiesIDs []int64) func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
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
				Field: "additional_cities",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidSpecialityID проверяет валидность выбранной специальности
var RuleValidSpecialityID = func(specialitiesIDs []int64) func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if !lo.Contains(specialitiesIDs, req.SpecialityID) {
			return false, dto.ValidationError{
				Text:  "Выбранной специальности не существует",
				Field: "speciality_id",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidSpecialitiesIDs проверяет валидность выбранных доп специальностей
var RuleValidSpecialitiesIDs = func(specialitiesIDs []int64) func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
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
				Field: "additional_specialities",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidSiteLink проверяет валидность сайта
var RuleValidSiteLink = func() func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if req.SiteLink == "" {
			return true, dto.ValidationError{}
		}

		restrictedDomains := []string{
			"t.me",
			"instagram.com",
			"wa.me",
			"dzen.ru",
			"youtube.com",
			"@",
			"vk.com",
		}

		for _, domain := range restrictedDomains {
			if strings.Contains(req.SiteLink, domain) {
				return false, dto.ValidationError{
					Text:  "Пожалуйста, укажите ссылку на сайт, а не на соц.сеть",
					Field: "site_link",
				}
			}
		}

		if !strings.HasPrefix(req.SiteLink, "http") {
			return false, dto.ValidationError{
				Text:  "Пожалуйста, укажите ссылку на сайт, ссылка должна содержать http",
				Field: "site_link",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidBirthDate проверяет валидность дня рождения
var RuleValidBirthDate = func() func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		birthDateStr := req.BirthDateString

		if birthDateStr == "" {
			return false, dto.ValidationError{
				Text:  "Дата рождения обязательна",
				Field: "birth_date",
			}
		}

		dateRegex := regexp.MustCompile(`^\d{2}\.\d{2}\.\d{4}$`)
		if !dateRegex.MatchString(birthDateStr) {
			return false, dto.ValidationError{
				Text:  "Неверный формат даты. ожидается ДД.ММ.ГГГГ",
				Field: "birth_date",
			}
		}

		birthDate, err := time.Parse("02.01.2006", birthDateStr)
		if err != nil {
			return false, dto.ValidationError{
				Text:  "Некорректная дата рождения",
				Field: "birth_date",
			}
		}

		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		if birthDate.After(today) {
			return false, dto.ValidationError{
				Text:  "Дата рождения не может быть в будущем",
				Field: "birth_date",
			}
		}

		maxAgeDate := today.AddDate(-120, 0, 0)
		if birthDate.Before(maxAgeDate) {
			return false, dto.ValidationError{
				Text:  "Дата рождения не должна быть старше 120 лет",
				Field: "birth_date",
			}
		}
		req.BirthDateTime = birthDate
		return true, dto.ValidationError{}
	}
}
