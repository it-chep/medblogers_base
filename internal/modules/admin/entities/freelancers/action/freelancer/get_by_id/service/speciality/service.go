package speciality

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
)

type Dal interface {
	GetSpecialities(ctx context.Context, freelancerID int64) ([]*speciality.Speciality, error)
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
	specialities, err := s.dal.GetSpecialities(ctx, freelancerDTO.ID)
	if err != nil {
		return nil, err
	}

	for _, spec := range specialities {
		if int64(spec.ID()) == freelancerDTO.Speciality.ID {
			freelancerDTO.Speciality = dto.Speciality{
				ID:   freelancerDTO.Speciality.ID,
				Name: spec.Name(),
			}
			continue
		}

		freelancerDTO.AdditionalSpecialities = append(freelancerDTO.AdditionalSpecialities, dto.Speciality{
			ID:   int64(spec.ID()),
			Name: spec.Name(),
		})
	}

	return freelancerDTO, nil
}
