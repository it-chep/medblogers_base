package rules

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/domain/speciality"
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
				Field: "instagramUsername",
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
				Field: "cityId",
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
				Field: "additionalCities",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidSpecialityID проверяет валидность выбранной специальности
var RuleValidSpecialityID = func(specialities []*speciality.Speciality) func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		var (
			validID   bool
			validSpec *speciality.Speciality
		)

		for _, spec := range specialities {
			if int64(spec.ID()) == req.SpecialityID {
				validID = true
				validSpec = spec
				break
			}
		}

		if !validID {
			return false, dto.ValidationError{
				Text:  "Выбранной специальности не существует",
				Field: "specialityId",
			}
		}

		if validSpec.IsAdditional() {
			return false, dto.ValidationError{
				Text:  "Эту специальность нельзя выбрать как основную",
				Field: "specialityId",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidSpecialitiesIDs проверяет валидность выбранных доп специальностей
var RuleValidSpecialitiesIDs = func(specialities []*speciality.Speciality) func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if len(req.AdditionalSpecialties) == 0 {
			return true, dto.ValidationError{}
		}

		validSpecialitiesMap := lo.SliceToMap(specialities, func(item *speciality.Speciality) (int64, *speciality.Speciality) {
			return int64(item.ID()), item
		})

		var invalidSpecialities []int64
		for _, id := range req.AdditionalSpecialties {
			_, exists := validSpecialitiesMap[id]
			if !exists {
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
					Field: "siteLink",
				}
			}
		}

		if !strings.HasPrefix(req.SiteLink, "http") {
			return false, dto.ValidationError{
				Text:  "Пожалуйста, укажите ссылку на сайт, ссылка должна содержать http",
				Field: "siteLink",
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
				Field: "birthDate",
			}
		}

		dateRegex := regexp.MustCompile(`^\d{2}\.\d{2}\.\d{4}$`)
		if !dateRegex.MatchString(birthDateStr) {
			return false, dto.ValidationError{
				Text:  "Неверный формат даты. ожидается ДД.ММ.ГГГГ",
				Field: "birthDate",
			}
		}

		birthDate, err := time.Parse("02.01.2006", birthDateStr)
		if err != nil {
			return false, dto.ValidationError{
				Text:  "Некорректная дата рождения",
				Field: "birthDate",
			}
		}

		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		if birthDate.After(today) {
			return false, dto.ValidationError{
				Text:  "Дата рождения не может быть в будущем",
				Field: "birthDate",
			}
		}

		maxAgeDate := today.AddDate(-120, 0, 0)
		if birthDate.Before(maxAgeDate) {
			return false, dto.ValidationError{
				Text:  "Дата рождения не должна быть старше 120 лет",
				Field: "birthDate",
			}
		}
		req.BirthDateTime = birthDate
		return true, dto.ValidationError{}
	}
}

// RuleValidInstagramLink валидирует ссылку на инстаграм
var RuleValidInstagramLink = func() func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if !strings.HasPrefix(req.InstagramUsername, "https://instagram.com") {
			req.InstagramUsername = fmt.Sprintf("https://instagram.com/%s", req.InstagramUsername)
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidTgChannelLink валидирует ссылку на тг канал
var RuleValidTgChannelLink = func() func(ctx context.Context, t *dto.CreateDoctorRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateDoctorRequest) (bool, error) {
		if !strings.HasPrefix(req.TelegramChannel, "https://t.me") {
			req.TelegramChannel = fmt.Sprintf("https://t.me/%s", req.TelegramChannel)
		}

		return true, dto.ValidationError{}
	}
}
