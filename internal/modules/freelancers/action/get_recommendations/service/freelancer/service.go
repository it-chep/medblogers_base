package freelancer

import (
	"context"
	domain "medblogers_base/internal/modules/freelancers/domain/freelancer_recommendation"
)

// Repository дал экшена
type Repository interface {
	GetRecommendations(ctx context.Context, freelancerID int64) (domain.FreelancerRecommendations, error)
}

// Service сервис по работе с фрилансерами
type Service struct {
	repository Repository
}

// New .
func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// GetRecommendations получение рекомендаций фрилансера
func (s *Service) GetRecommendations(ctx context.Context, freelancerID int64) (domain.FreelancerRecommendations, error) {
	return s.repository.GetRecommendations(ctx, freelancerID)
}
