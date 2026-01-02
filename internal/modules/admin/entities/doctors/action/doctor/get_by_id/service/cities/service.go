package cities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
)

type Dal interface {
	GetAdditionalCities(ctx context.Context, doctorID int64) ([]*city.City, error)
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

	cities, err := s.dal.GetAdditionalCities(ctx, docDTO.ID)
	if err != nil {
		return nil, err
	}

	for _, c := range cities {
		if int64(c.ID()) == docDTO.MainCity.ID {
			docDTO.MainCity = dto.City{
				ID:   docDTO.MainCity.ID,
				Name: c.Name(),
			}
			continue
		}

		docDTO.AdditionalCities = append(docDTO.AdditionalCities, dto.City{
			ID:   int64(c.ID()),
			Name: c.Name(),
		})
	}

	return docDTO, nil
}
