package settings

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/async"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . CityStorage,SpecialityStorage,DoctorsStorage

// CityStorage .
type CityStorage interface {
	GetCities(ctx context.Context) ([]*city.City, error)
}

// SpecialityStorage .
type SpecialityStorage interface {
	GetSpecialities(ctx context.Context) ([]*speciality.Speciality, error)
}

// DoctorsStorage .
type DoctorsStorage interface {
	GetDoctorsCount(ctx context.Context) (int64, error)
}

// Service .
type Service struct {
	cityStorage     CityStorage
	specialityStory SpecialityStorage
	doctorsStorage  DoctorsStorage
}

// NewSettingsService .
func NewSettingsService(cityStorage CityStorage, specialityStory SpecialityStorage, doctorsStorage DoctorsStorage) *Service {
	return &Service{
		cityStorage:     cityStorage,
		specialityStory: specialityStory,
		doctorsStorage:  doctorsStorage,
	}
}

func (s *Service) GetSettings(ctx context.Context) error {
	// Делаем обычную группу, чтобы если 1 из фильров не работает, он не ломал нам весь сайт
	g := async.NewGroup()

	// получение городов
	g.Go(func() {
		cities, err := s.cityStorage.GetCities(ctx)
		if err != nil {
			//	todo log
		}
	})

	// получение специальностей
	g.Go(func() {
		specialities, err := s.specialityStory.GetSpecialities(ctx)
		if err != nil {
			//	todo log
		}
	})

	// получение количества докторов
	g.Go(func() {
		doctorsCount, err := s.doctorsStorage.GetDoctorsCount(ctx)
		if err != nil {
			//	todo log
		}
	})

	g.Wait()

	return nil
}
