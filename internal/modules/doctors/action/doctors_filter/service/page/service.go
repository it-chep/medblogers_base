package page

import (
	"math"
	consts "medblogers_base/internal/dto"
)

type Service struct {
}

func New() *Service {
	return &Service{}
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
