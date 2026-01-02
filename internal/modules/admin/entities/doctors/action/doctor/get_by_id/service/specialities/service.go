package specialities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
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

func (s *Service) Enrich(ctx context.Context, docDTO *dto.DoctorDTO) (*dto.DoctorDTO, error) {

	specialities, err := s.dal.GetAdditionalSpecialities(ctx, docDTO.ID)
	if err != nil {
		return nil, err
	}

	for _, spec := range specialities {
		if int64(spec.ID()) == docDTO.MainSpeciality.ID {
			docDTO.MainSpeciality = dto.Speciality{
				ID:   docDTO.MainSpeciality.ID,
				Name: spec.Name(),
			}
			continue
		}

		docDTO.AdditionalSpecialities = append(docDTO.AdditionalSpecialities, dto.Speciality{
			ID:   int64(spec.ID()),
			Name: spec.Name(),
		})
	}

	return docDTO, nil
}
