package cooperation_type

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
)

type Dal interface {
	GetCooperationType(ctx context.Context, cooperationTypeID int64) (*doctor.CooperationType, error)
}

type Service struct {
	dal Dal
}

func New(dal Dal) *Service {
	return &Service{
		dal: dal,
	}
}

func (s *Service) Enrich(ctx context.Context, docDTO *dto.DoctorDTO) (*dto.DoctorDTO, error) {

	cooperationType, err := s.dal.GetCooperationType(ctx, docDTO.CooperationType.ID)
	if err != nil {
		return nil, err
	}

	docDTO.CooperationType.Name = cooperationType.Name()

	return docDTO, nil
}
