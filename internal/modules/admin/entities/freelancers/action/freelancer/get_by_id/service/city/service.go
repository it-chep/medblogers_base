package city

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
)

type Dal interface {
	GetCities(ctx context.Context, freelancerID int64) ([]*city.City, error)
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
	cities, err := s.dal.GetCities(ctx, freelancerDTO.ID)
	if err != nil {
		return nil, err
	}

	for _, c := range cities {
		if int64(c.ID()) == freelancerDTO.City.ID {
			freelancerDTO.City = dto.City{
				ID:   freelancerDTO.City.ID,
				Name: c.Name(),
			}
			continue
		}

		freelancerDTO.AdditionalCities = append(freelancerDTO.AdditionalCities, dto.City{
			ID:   int64(c.ID()),
			Name: c.Name(),
		})
	}

	return freelancerDTO, nil
}
