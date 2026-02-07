package speciality

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
)

type Dal interface {
	GetSpeciality(ctx context.Context, specialityID int64) (*speciality.Speciality, error)
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
	specialityDomain, err := s.dal.GetSpeciality(ctx, freelancerDTO.ID)
	if err != nil {
		return nil, err
	}

	freelancerDTO.Speciality = dto.Speciality{
		ID:   int64(specialityDomain.ID()),
		Name: specialityDomain.Name(),
	}

	return freelancerDTO, nil
}
