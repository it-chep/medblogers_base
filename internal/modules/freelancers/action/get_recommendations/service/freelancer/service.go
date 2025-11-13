package freelancer

import (
	"context"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	domain "medblogers_base/internal/modules/freelancers/domain/freelancer_recommendation"
)

// Repository дал экшена
type Repository interface {
	GetRecommendations(ctx context.Context, freelancerID int64) (domain.FreelancerRecommendations, error)
}

type FreelancerDal interface {
	GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error)
}

// Service сервис по работе с фрилансерами
type Service struct {
	repository    Repository
	freelancerDal FreelancerDal
}

// New .
func New(repository Repository, freelancerDal FreelancerDal) *Service {
	return &Service{
		repository:    repository,
		freelancerDal: freelancerDal,
	}
}

// GetFreelancerInfo получение инфы о фрилансере
func (s *Service) GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error) {
	return s.freelancerDal.GetFreelancerInfo(ctx, slug)
}

// GetRecommendations получение рекомендаций фрилансера
func (s *Service) GetRecommendations(ctx context.Context, freelancerID int64) (domain.FreelancerRecommendations, error) {
	return s.repository.GetRecommendations(ctx, freelancerID)
}
