package subscribers

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/get_pages_count/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"

	"github.com/samber/lo"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . SubscribersGetter

// SubscribersGetter получение информации о подписчиках
type SubscribersGetter interface {
	GetDoctorsByFilter(ctx context.Context, request indto.GetDoctorsByFilterRequest) (indto.GetDoctorsByFilterResponse, error)
	GetDoctorsByFilterWithIDs(ctx context.Context, request indto.GetDoctorsByFilterRequest, doctorsIDs []int64) (indto.GetDoctorsByFilterResponse, error)
}

type Service struct {
	getter SubscribersGetter
}

func New(getter SubscribersGetter) *Service {
	return &Service{
		getter: getter,
	}
}

func (s *Service) FilterDoctorsBySubscribersWithDoctorsIDs(ctx context.Context, filter dto.Filter, doctorsIDs []int64) (int64, error) {
	response, err := s.getter.GetDoctorsByFilterWithIDs(
		ctx, indto.GetDoctorsByFilterRequest{
			MinSubscribers: filter.MinSubscribers,
			MaxSubscribers: filter.MaxSubscribers,
			SocialMedia: lo.Map(filter.SocialMedia, func(socialMedia string, index int) indto.SocialMedia {
				return indto.NewSocialMedia(socialMedia)
			}),
		},
		doctorsIDs,
	)
	if err != nil {
		return 0, err
	}

	return response.DoctorsCount, nil
}

// FilterDoctorsBySubscribers - фильтрация по количеству подписчиков
func (s *Service) FilterDoctorsBySubscribers(ctx context.Context, filter dto.Filter) (int64, error) {
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
		return 0, err
	}

	return response.DoctorsCount, nil
}
