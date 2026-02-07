package cooperation_type

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
)

type Dal interface {
	GetCooperationType(ctx context.Context, cooperationTypeID int64) (*freelancer.CooperationType, error)
}

type Service struct {
	dal Dal
}

func New(dal Dal) *Service {
	return &Service{
		dal: dal,
	}
}

func (s *Service) Enrich(ctx context.Context, frlnsr *dto.FreelancerDTO) (*dto.FreelancerDTO, error) {

	cooperationType, err := s.dal.GetCooperationType(ctx, frlnsr.CooperationType.ID)
	if err != nil {
		return nil, err
	}

	frlnsr.CooperationType.Name = cooperationType.Name()

	return frlnsr, nil
}
