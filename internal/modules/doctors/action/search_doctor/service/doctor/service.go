package doctor

import (
	"context"

	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/city"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/doctor"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/speciality"
	"github.com/it-chep/medblogers_base/internal/pkg/async"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . CityStorage,SpecialityStorage,DoctorsStorage

// CityStorage .
type CityStorage interface {
	SearchCities(ctx context.Context, query string) ([]*city.City, error)
}

// SpecialityStorage .
type SpecialityStorage interface {
	SearchSpecialities(ctx context.Context, query string) ([]*speciality.Speciality, error)
}

// DoctorsStorage .
type DoctorsStorage interface {
	SearchDoctors(ctx context.Context, query string) ([]*doctor.Doctor, error)
}

// Service .
type Service struct {
	cityStorage     CityStorage
	specialityStory SpecialityStorage
	doctorsStorage  DoctorsStorage
}

// NewSearchService .
func NewSearchService(cityStorage CityStorage, specialityStory SpecialityStorage, doctorsStorage DoctorsStorage) *Service {
	return &Service{
		cityStorage:     cityStorage,
		specialityStory: specialityStory,
		doctorsStorage:  doctorsStorage,
	}
}

func (s *Service) Search(ctx context.Context, query string) error {
	// Делаем обычную группу, чтобы если 1 из фильров не работает, он не ломал нам весь сайт
	g := async.NewGroup()

	// получение городов
	g.Go(func() {
		cities, err := s.cityStorage.SearchCities(ctx, query)
		if err != nil {
			//	todo log
		}
	})

	// получение специальностей
	g.Go(func() {
		specialities, err := s.specialityStory.SearchSpecialities(ctx, query)
		if err != nil {
			//	todo log
		}
	})

	// получение количества докторов
	g.Go(func() {
		doctors, err := s.doctorsStorage.SearchDoctors(ctx, query)
		if err != nil {
			//	todo log
		}
	})

	g.Wait()
}
