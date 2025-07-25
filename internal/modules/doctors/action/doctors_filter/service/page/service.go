package page

import (
	"context"
	"math"
	consts "medblogers_base/internal/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	GetDoctorsCountByFilter(ctx context.Context, filter dto.Filter) (int64, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

// GetPagesCount получение количества страниц для пагинации
func (s *Service) GetPagesCount(ctx context.Context, filter dto.Filter) (int64, error) {
	doctorsCount, err := s.storage.GetDoctorsCountByFilter(ctx, filter)
	if err != nil {
		return 1, err
	}

	return s.countPages(doctorsCount), nil
}

// GetPagesCountBySubscribersFilter получение количества страниц для пагинации фильтра с подписчиками
func (s *Service) GetPagesCountBySubscribersFilter(doctorsFromSubsCount int64) int64 {
	return s.countPages(doctorsFromSubsCount)
}

func (s *Service) countPages(doctorsCount int64) int64 {
	pagesCount := int64(math.Ceil(float64(doctorsCount) / float64(consts.LimitDoctorsOnPage)))
	if pagesCount < 1 {
		pagesCount = 1
	}

	return pagesCount
}
