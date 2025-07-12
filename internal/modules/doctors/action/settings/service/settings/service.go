package settings

import (
	"context"

	"github.com/it-chep/medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/city"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/speciality"
	"github.com/it-chep/medblogers_base/internal/pkg/async"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . CityStorage,SpecialityStorage,DoctorsStorage,SubscribersGetter

// CityStorage .
type CityStorage interface {
	GetCitiesWithDoctorsCount(ctx context.Context) ([]*city.City, error)
}

// SpecialityStorage .
type SpecialityStorage interface {
	GetSpecialitiesWithDoctorsCount(ctx context.Context) ([]*speciality.Speciality, error)
}

// DoctorsStorage .
type DoctorsStorage interface {
	GetDoctorsCount(ctx context.Context) (int64, error)
}

type SubscribersGetter interface {
	GetAllSubscribersInfo(ctx context.Context) (indto.GetAllSubscribersInfoResponse, error)
}

// Service .
type Service struct {
	cityStorage       CityStorage
	specialityStory   SpecialityStorage
	doctorsStorage    DoctorsStorage
	subscribersGetter SubscribersGetter
}

// NewSettingsService .
func NewSettingsService(cityStorage CityStorage, specialityStory SpecialityStorage, doctorsStorage DoctorsStorage, subscribersGetter SubscribersGetter) *Service {
	return &Service{
		cityStorage:       cityStorage,
		specialityStory:   specialityStory,
		doctorsStorage:    doctorsStorage,
		subscribersGetter: subscribersGetter,
	}
}

func (s *Service) GetSettings(ctx context.Context) error {
	// Делаем обычную группу, чтобы если 1 из фильров не работает, он не ломал нам весь сайт
	g := async.NewGroup()

	// получение городов
	g.Go(func() {
		_, err := s.cityStorage.GetCitiesWithDoctorsCount(ctx)
		if err != nil {
			//	todo log
		}
	})

	// получение специальностей
	g.Go(func() {
		_, err := s.specialityStory.GetSpecialitiesWithDoctorsCount(ctx)
		if err != nil {
			//	todo log
		}
	})

	// получение количества докторов
	g.Go(func() {
		_, err := s.doctorsStorage.GetDoctorsCount(ctx)
		if err != nil {
			//	todo log
		}
	})

	g.Go(func() {
		_, err := s.subscribersGetter.GetAllSubscribersInfo(ctx)
		if err != nil {
			//	todo log
		}
	})

	g.Wait()

	return nil
}
