package specialities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/speciality"
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

func (s *Service) Enrich(ctx context.Context, docDTO *dto.DoctorDTO) (*dto.DoctorDTO, error) {

	spec, err := s.dal.GetSpeciality(ctx, docDTO.MainSpeciality.ID)
	if err != nil {
		return nil, err
	}

	docDTO.MainSpeciality = dto.Speciality{
		ID:   int64(spec.ID()),
		Name: spec.Name(),
	}

	return docDTO, nil
}
