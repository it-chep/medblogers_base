package cities

import (
	"context"
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

func (s *Service) GetAdditionalCities(ctx context.Context, doctorID int64) ([]*city.City, error) {
	cities, err := s.dal.GetAdditionalCities(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
