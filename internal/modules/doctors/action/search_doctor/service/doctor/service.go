package doctor

import (
	"context"

	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/async"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . SearchStorage

// SearchStorage .
type SearchStorage interface {
	SearchCities(ctx context.Context, query string) ([]*city.City, error)
	SearchSpecialities(ctx context.Context, query string) ([]*speciality.Speciality, error)
	SearchDoctors(ctx context.Context, query string) ([]*doctor.Doctor, error)
}

// Service .
type Service struct {
	searchStorage SearchStorage
}

// NewSearchService .
func NewSearchService(searchStorage SearchStorage) *Service {
	return &Service{
		searchStorage: searchStorage,
	}
}

func (s *Service) Search(ctx context.Context, query string) error {
	// Делаем обычную группу, чтобы если 1 из фильров не работает, он не ломал нам весь сайт
	g := async.NewGroup()

	// получение городов
	g.Go(func() {
		_, err := s.searchStorage.SearchCities(ctx, query)
		if err != nil {
			//	todo log
		}
	})

	// получение специальностей
	g.Go(func() {
		_, err := s.searchStorage.SearchSpecialities(ctx, query)
		if err != nil {
			//	todo log
		}
	})

	// получение количества докторов
	g.Go(func() {
		_, err := s.searchStorage.SearchDoctors(ctx, query)
		if err != nil {
			//	todo log
		}
	})

	g.Wait()

	return nil
}
