package cities

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
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

func (s *Service) Enrich(ctx context.Context, docDTO *dto.DoctorDTO) (*dto.DoctorDTO, error) {
	cit, err := s.dal.GetCity(ctx, docDTO.MainCity.ID)
	if err != nil {
		return nil, err
	}

	docDTO.MainCity = dto.City{
		ID:   int64(cit.ID()),
		Name: cit.Name(),
	}

	return docDTO, nil
}
