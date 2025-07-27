package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
	"medblogers_base/internal/pkg/logger"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	CountFilterDoctors(ctx context.Context, filter dto.Filter) (int64, error)
	CountByFilterAndIDs(ctx context.Context, filter dto.Filter, doctorIDs []int64) (int64, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

// GetDoctorsByFilter - фильтрация докторов по полям в базе
func (s *Service) GetDoctorsByFilter(ctx context.Context, filter dto.Filter) (int64, error) {
	logger.Message(ctx, "[PreliminaryFilterCount][Service] Получение количества докторов по фильтрам")

	doctorsCount, err := s.storage.CountFilterDoctors(ctx, filter)
	if err != nil {
		return 0, err
	}

	return doctorsCount, nil
}

// GetDoctorsByFilterAndIDs - фильтрация докторов по фильтрам и ID из subscribers
func (s *Service) GetDoctorsByFilterAndIDs(ctx context.Context, filter dto.Filter, ids []int64) (int64, error) {
	logger.Message(ctx, "[PreliminaryFilterCount][Service] Получение количества докторов по фильтрам")

	doctorsCount, err := s.storage.CountByFilterAndIDs(ctx, filter, ids)
	if err != nil {
		return 0, err
	}

	return doctorsCount, nil
}
