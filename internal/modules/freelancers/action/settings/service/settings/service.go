package settings

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/settings/dto"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
)

type CityStorage interface {
	GetCitiesWithFreelancersCount(ctx context.Context) ([]*city.City, error)
}

type SpecialityStorage interface {
	GetSpecialitiesWithFreelancersCount(ctx context.Context) ([]*speciality.Speciality, error)
}

type SocietyStorage interface {
	GetCitiesWithFreelancersCount(ctx context.Context) ([]*social_network.SocialNetwork, error)
}

type CategoryStorage interface {
	GetPriceCategoriesInfo(ctx context.Context) ([]dto.PriceCategory, error)
}

type Service struct {
	cityStorage       CityStorage
	specialityStorage SpecialityStorage
	societyStorage    SocietyStorage
	categoryStorage   CategoryStorage
}

func NewSettingsService(
	cityStorage CityStorage, specialityStorage SpecialityStorage,
	societyStorage SocietyStorage, categoryStorage CategoryStorage,
) *Service {
	return &Service{
		cityStorage:       cityStorage,
		specialityStorage: specialityStorage,
		societyStorage:    societyStorage,
		categoryStorage:   categoryStorage,
	}
}

func (s *Service) GetSettings(ctx context.Context) (*dto.Settings, error) {
	var (
		cities       []*city.City
		specialities []*speciality.Speciality
		categories   []dto.PriceCategory
		networks     []*social_network.SocialNetwork
	)
	// Делаем обычную группу, чтобы если 1 из фильров не работает, он не ломал нам весь сайт
	g := async.NewGroup()

	// получение городов
	g.Go(func() {
		gCities, gErr := s.cityStorage.GetCitiesWithFreelancersCount(ctx)
		if gErr != nil {
			logger.Error(ctx, "[Settings] Ошибка при получении городов", gErr)
		}
		cities = gCities
	})

	// получение специальностей
	g.Go(func() {
		gSpecialities, gErr := s.specialityStorage.GetSpecialitiesWithFreelancersCount(ctx)
		if gErr != nil {
			logger.Error(ctx, "[Settings] Ошибка при получении специальностей", gErr)
		}
		specialities = gSpecialities
	})

	// соц сети
	g.Go(func() {
		gNetworks, gErr := s.societyStorage.GetCitiesWithFreelancersCount(ctx)
		if gErr != nil {
			logger.Error(ctx, "[Settings] Ошибка при получении соц сетей", gErr)
		}
		networks = gNetworks
	})

	// Ценовые категории
	g.Go(func() {
		gCategories, gErr := s.categoryStorage.GetPriceCategoriesInfo(ctx)
		if gErr != nil {
			logger.Error(ctx, "[Settings] Ошибка при получении ценовых категорий", gErr)
		}
		categories = gCategories
	})

	g.Wait()

	settings := dto.NewSettings(cities, specialities, categories, networks)

	return settings, nil
}
