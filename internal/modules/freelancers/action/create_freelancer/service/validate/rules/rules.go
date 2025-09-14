package rules

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
	"regexp"
	"strings"
	"unicode/utf8"
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

// RuleValidSocialNetworksIDs проверяет валидность выбранных соц сетей
var RuleValidSocialNetworksIDs = func(networksIDs []int64) func(ctx context.Context, t *dto.CreateRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateRequest) (bool, error) {
		if len(req.SocialNetworks) == 0 {
			return true, dto.ValidationError{}
		}

		validNetworksMap := make(map[int64]struct{})
		for _, id := range networksIDs {
			validNetworksMap[id] = struct{}{}
		}

		var invalidNetworks []int64
		for _, id := range req.SocialNetworks {
			if _, exists := validNetworksMap[id]; !exists {
				invalidNetworks = append(invalidNetworks, id)
			}
		}

		if len(invalidNetworks) > 0 {
			return false, dto.ValidationError{
				Text:  "Выбранных соц сетей не существует",
				Field: "socialNetworks",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidPortfolioLink валидирует ссылку на портфолио
var RuleValidPortfolioLink = func() func(ctx context.Context, t *dto.CreateRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateRequest) (bool, error) {
		if utf8.RuneCountInString(req.PortfolioLink) == 0 {
			return true, dto.ValidationError{}
		}

		if !strings.HasPrefix(req.PortfolioLink, "https://") {
			return false, dto.ValidationError{
				Text:  "Необходимо указать ссылку на портфолио",
				Field: "portfolioLink",
			}
		}

		return true, dto.ValidationError{}
	}
}

// RuleValidTgUsername валидирует тг
var RuleValidTgUsername = func() func(ctx context.Context, t *dto.CreateRequest) (bool, error) {
	return func(_ context.Context, req *dto.CreateRequest) (bool, error) {
		if utf8.RuneCountInString(req.TgUsername) == 0 {
			return false, dto.ValidationError{
				Text:  "Обязательное поле",
				Field: "telegramUsername",
			}
		}

		link := strings.TrimSpace(req.TgUsername)

		// Регулярное выражение для проверки Telegram username/ссылки
		tgPattern := `^(?:@|(?:https?://)?t\.me/)?([a-zA-Z][a-zA-Z0-9_]{4,31})$`
		re := regexp.MustCompile(tgPattern)

		matches := re.FindStringSubmatch(link)
		if matches == nil {
			return false, dto.ValidationError{
				Text:  "Укажите валидный Telegram username или ссылку (например: @username, t.me/username)",
				Field: "telegramUsername",
			}
		}

		return true, nil
	}
}
