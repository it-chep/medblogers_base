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

// GetPagesCount получение количества страниц для пагинации
func (s *Service) GetPagesCount(freelancersCount int64) int64 {
	return s.countPages(freelancersCount)
}

func (s *Service) countPages(freelancersCount int64) int64 {
	pagesCount := int64(math.Ceil(float64(freelancersCount) / float64(consts.LimitDoctorsOnPage)))
	if pagesCount < 1 {
		pagesCount = 1
	}

	return pagesCount
}
