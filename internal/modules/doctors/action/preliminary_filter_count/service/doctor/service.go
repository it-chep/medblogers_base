package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/logger"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	FilterDoctors(ctx context.Context, filter dto.Filter) (map[doctor.MedblogersID]*doctor.Doctor, error)
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
func (s *Service) GetDoctorsByFilter(ctx context.Context, filter dto.Filter) ([]int64, error) {
	logger.Message(ctx, "[PreliminaryFilterCount][Service] Получение количества докторов по фильтрам")

	doctorsMap, err := s.storage.FilterDoctors(ctx, filter)
	if err != nil {
		return nil, err
	}

	doctorIDs := make([]int64, 0, len(doctorsMap))
	for _, doc := range doctorsMap {
		doctorIDs = append(doctorIDs, int64(doc.GetID()))
	}

	return doctorIDs, nil
}
