package freelancer

import (
	"context"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
)

// Storage .
type Storage interface {
	GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error)
}

// Service сервис получения данных о фрилансере
type Service struct {
	storage Storage
}

// New к-ор
func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

// GetFreelancerInfo получение информации о фрилансере
func (s *Service) GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error) {
	return s.storage.GetFreelancerInfo(ctx, slug)
}
