package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage,CommonDal

type CommonDal interface {
	GetDoctorsCount(ctx context.Context) (int64, error)
}

type Storage interface {
	FilterDoctors(ctx context.Context, filter dto.Filter) ([]int64, error)
}

type Service struct {
	commonDal CommonDal
	storage   Storage
}

func New(storage Storage, commonDal CommonDal) *Service {
	return &Service{
		commonDal: commonDal,
		storage:   storage,
	}
}

func (s *Service) GetDoctorsCount(ctx context.Context) (int64, error) {
	return s.commonDal.GetDoctorsCount(ctx)
}

// GetDoctorsIDsByFilter - фильтрация докторов по полям в базе
func (s *Service) GetDoctorsIDsByFilter(ctx context.Context, filter dto.Filter) ([]int64, error) {
	orderedIDs, err := s.storage.FilterDoctors(ctx, filter)
	if err != nil {
		return nil, err
	}

	return orderedIDs, nil
}
