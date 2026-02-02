package city

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
)

type Dal interface {
	GetCity(ctx context.Context, cityID int64) (*city.City, error)
}
type Service struct {
	dal Dal
}

func New(dal Dal) *Service {
	return &Service{
		dal: dal,
	}
}

func (s *Service) Enrich(ctx context.Context, freelancerDTO *dto.FreelancerDTO) (*dto.FreelancerDTO, error) {
	cityDomain, err := s.dal.GetCity(ctx, freelancerDTO.ID)
	if err != nil {
		return nil, err
	}

	freelancerDTO.City = dto.City{
		ID:   int64(cityDomain.ID()),
		Name: cityDomain.Name(),
	}

	return freelancerDTO, nil
}
