package freelancers

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count/dto"
)

type Storage interface {
	FilterFreelancers(ctx context.Context, filter dto.Filter) (int64, error)
}

// Service .
type Service struct {
	repository Storage
}

// NewService .
func NewService(repository Storage) *Service {
	return &Service{
		repository: repository,
	}
}

// GetFreelancersCount селект количества врачей из базы
func (s *Service) GetFreelancersCount(ctx context.Context, filter dto.Filter) (int64, error) {
	return s.repository.FilterFreelancers(ctx, filter)
}
