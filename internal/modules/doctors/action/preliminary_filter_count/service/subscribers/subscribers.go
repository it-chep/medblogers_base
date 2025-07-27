package subscribers

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"

	"github.com/samber/lo"

	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . SubscribersGetter

// SubscribersGetter получение информации о подписчиках
type SubscribersGetter interface {
	GetDoctorsByFilter(ctx context.Context, request indto.GetDoctorsByFilterRequest) (map[int64]indto.GetDoctorsByFilterResponse, error)
}

type Service struct {
	getter SubscribersGetter
}

func New(getter SubscribersGetter) *Service {
	return &Service{
		getter: getter,
	}
}

// FilterDoctorsBySubscribers - фильтрация по количеству подписчиков
func (s *Service) FilterDoctorsBySubscribers(ctx context.Context, filter dto.Filter) ([]int64, error) {
	response, err := s.getter.GetDoctorsByFilter(
		ctx, indto.GetDoctorsByFilterRequest{
			MinSubscribers: filter.MinSubscribers,
			MaxSubscribers: filter.MaxSubscribers,
			SocialMedia: lo.Map(filter.SocialMedia, func(socialMedia string, index int) indto.SocialMedia {
				return indto.NewSocialMedia(socialMedia)
			}),
		},
	)
	if err != nil {
		return nil, err
	}

	return lo.Keys(response), nil
}
