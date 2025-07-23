package rules

import (
	"context"
	"errors"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/pkg/spec"
	"regexp"
	"strings"
	"time"
)

// RuleAtLeastOneSocialMedia проверяет что у пользователя указана хотя бы 1 соцсеть
var RuleAtLeastOneSocialMedia = func() spec.Specification[*dto.CreateDoctorRequest] {
	return spec.NewSpecification(func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if req.InstagramUsername == "" &&
			req.VKUsername == "" &&
			req.TelegramUsername == "" &&
			req.TelegramChannel == "" &&
			req.DzenUsername == "" &&
			req.YoutubeUsername == "" {
			return false, errors.New("обязательно нужно указать хотя бы 1 вашу соцсеть")
		}
		return true, nil
	})
}

// RuleValidCityID проверяет валидность выбранного города
var RuleValidCityID = func(citiesIDs []int64) spec.Specification[*dto.CreateDoctorRequest] {
	return spec.NewSpecification(func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if !lo.Contains(citiesIDs, req.CityID) {
			return false, errors.New("Выбранного города не существует")
		}

		return true, nil
	})
}

// RuleValidAdditionalCitiesIDs проверяет валидность выбранных доп города
var RuleValidAdditionalCitiesIDs = func(citiesIDs []int64) spec.Specification[*dto.CreateDoctorRequest] {
	return spec.NewSpecification(func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if len(req.AdditionalCities) == 0 {
			return true, nil
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
			return false, errors.New("содержатся недопустимые ID дополнительных городов")
		}

		return true, nil
	})
}

// RuleValidSpecialityID проверяет валидность выбранной специальности
var RuleValidSpecialityID = func(specialitiesIDs []int64) spec.Specification[*dto.CreateDoctorRequest] {
	return spec.NewSpecification(func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if !lo.Contains(specialitiesIDs, req.SpecialityID) {
			return false, errors.New("Выбранной специальности не существует")
		}

		return true, nil
	})
}

// RuleValidSpecialitiesIDs проверяет валидность выбранных доп специальностей
var RuleValidSpecialitiesIDs = func(specialitiesIDs []int64) spec.Specification[*dto.CreateDoctorRequest] {
	return spec.NewSpecification(func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if len(req.AdditionalCities) == 0 {
			return true, nil
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
			return false, errors.New("содержатся недопустимые ID дополнительных специальностей")
		}

		return true, nil
	})
}

// RuleValidSiteLink проверяет валидность сайта
var RuleValidSiteLink = func() spec.Specification[*dto.CreateDoctorRequest] {
	return spec.NewSpecification(func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if !strings.HasPrefix(req.SiteLink, "http") {
			return false, errors.New("пожалуйста, укажите ссылку на сайт, ссылка должна содержать http")
		}

		return true, nil
	})
}

// RuleValidBirthDate проверяет валидность дня рождения
var RuleValidBirthDate = func() spec.Specification[*dto.CreateDoctorRequest] {
	return spec.NewSpecification(func(ctx context.Context, doc *dto.CreateDoctorRequest) (bool, error) {
		birthDateStr := doc.BirthDate

		if birthDateStr == "" {
			return false, errors.New("дата рождения обязательна")
		}

		dateRegex := regexp.MustCompile(`^\d{2}\.\d{2}\.\d{4}$`)
		if !dateRegex.MatchString(birthDateStr) {
			return false, errors.New("неверный формат даты. ожидается ДД.ММ.ГГГГ")
		}

		birthDate, err := time.Parse("02.01.2006", birthDateStr)
		if err != nil {
			return false, errors.New("некорректная дата рождения")
		}

		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		if birthDate.After(today) {
			return false, errors.New("дата рождения не может быть в будущем")
		}

		maxAgeDate := today.AddDate(-120, 0, 0)
		if birthDate.Before(maxAgeDate) {
			return false, errors.New("дата рождения не должна быть старше 120 лет")
		}

		return true, nil
	})
}
