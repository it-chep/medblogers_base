package specialities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
)

type Dal interface {
	GetAdditionalSpecialities(ctx context.Context, doctorID int64) ([]*speciality.Speciality, error)
}

type Service struct {
	dal Dal
}

func New(dal Dal) *Service {
	return &Service{
		dal: dal,
	}
}

func (s *Service) GetAdditionalSpecialities(ctx context.Context, doctorID int64) ([]*speciality.Speciality, error) {
	specialities, err := s.dal.GetAdditionalSpecialities(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	return specialities, nil
}
