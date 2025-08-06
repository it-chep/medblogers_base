package city

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/pkg/logger"
)

// SearchStorage .
type SearchStorage interface {
	SearchCities(ctx context.Context, query string) ([]*city.City, error)
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

// SearchCities поиск городов
func (svc *Service) SearchCities(ctx context.Context, query string) ([]*city.City, error) {
	logger.Message(ctx, "[Search] Поиск городов ...")
	return svc.searchStorage.SearchCities(ctx, query)
}
