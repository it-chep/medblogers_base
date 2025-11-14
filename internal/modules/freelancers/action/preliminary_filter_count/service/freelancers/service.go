package freelancers

import (
	"context"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
)

type Storage interface {
	FreelancersCountByFilter(ctx context.Context, filter freelancer.Filter) (int64, error)
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
func (s *Service) GetFreelancersCount(ctx context.Context, filter freelancer.Filter) (int64, error) {
	return s.repository.FreelancersCountByFilter(ctx, filter)
}
