package settings

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/settings/dto"
	"medblogers_base/internal/pkg/logger"

	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/async"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . CityStorage,SpecialityStorage,SubscribersGetter

// CityStorage .
type CityStorage interface {
	GetCitiesWithDoctorsCount(ctx context.Context) ([]*city.City, error)
}

// SpecialityStorage .
type SpecialityStorage interface {
	GetSpecialitiesWithDoctorsCount(ctx context.Context) ([]*speciality.Speciality, error)
}

// SubscribersGetter .
type SubscribersGetter interface {
	GetFilterInfo(ctx context.Context) ([]indto.FilterInfoResponse, error)
}

// Service .
type Service struct {
	cityStorage       CityStorage
	specialityStory   SpecialityStorage
	subscribersGetter SubscribersGetter
}

// NewSettingsService .
func NewSettingsService(cityStorage CityStorage, specialityStory SpecialityStorage, subscribersGetter SubscribersGetter) *Service {
	return &Service{
		cityStorage:       cityStorage,
		specialityStory:   specialityStory,
		subscribersGetter: subscribersGetter,
	}
}

// GetSettings - получение настроек для главной страницы
func (s *Service) GetSettings(ctx context.Context) (_ *dto.Settings, err error) {
	var (
		cities         []*city.City
		specialities   []*speciality.Speciality
		enabledFilters []indto.FilterInfoResponse
	)
	// Делаем обычную группу, чтобы если 1 из фильров не работает, он не ломал нам весь сайт
	g := async.NewGroup()

	// получение городов
	g.Go(func() {
		gCities, gErr := s.cityStorage.GetCitiesWithDoctorsCount(ctx)
		if gErr != nil {
			logger.Error(ctx, "[Settings] Ошибка при получении городов", err)
		}
		cities = gCities
	})

	// получение специальностей
	g.Go(func() {
		gSpecialities, gErr := s.specialityStory.GetSpecialitiesWithDoctorsCount(ctx)
		if gErr != nil {
			logger.Error(ctx, "[Settings] Ошибка при получении специальностей", err)
		}
		specialities = gSpecialities
	})

	g.Go(func() {
		gEnabledFilters, gErr := s.subscribersGetter.GetFilterInfo(ctx)
		if gErr != nil {
			logger.Error(ctx, "[Settings] Ошибка при получении специальностей", err)
		}
		enabledFilters = gEnabledFilters
	})

	g.Wait()

	settings := dto.NewSettings(cities, specialities, enabledFilters)

	return settings, nil
}
