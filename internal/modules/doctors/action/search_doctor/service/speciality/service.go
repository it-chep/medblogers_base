package speciality

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/logger"
)

// SearchStorage .
type SearchStorage interface {
	SearchSpecialities(ctx context.Context, query string) ([]*speciality.Speciality, error)
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

// SearchSpecialities поиск специальностей
func (svc *Service) SearchSpecialities(ctx context.Context, query string) ([]*speciality.Speciality, error) {
	logger.Message(ctx, "[Search] Поиск специальностей ...")
	return svc.searchStorage.SearchSpecialities(ctx, query)
}
