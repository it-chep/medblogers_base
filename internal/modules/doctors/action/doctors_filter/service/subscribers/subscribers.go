package subscribers

import (
	"context"

	"github.com/samber/lo"

	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
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
func (s *Service) FilterDoctorsBySubscribers(ctx context.Context, filter *dto.Filter) (map[int64]dto.SubscribersInfoDTO, error) {
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

	result := make(map[int64]dto.SubscribersInfoDTO, len(response))
	for _, doctor := range response {
		result[doctor.DoctorID] = dto.SubscribersInfoDTO{
			InstSubsCount:     doctor.InstSubsCount,
			InstSubsCountText: doctor.InstSubsCountText,
			TgSubsCount:       doctor.TgSubsCount,
			TgSubsCountText:   doctor.TgSubsCountText,
		}
	}

	return result, nil
}

func (s *Service) EnrichSubscribers(doctorsMap map[int64]dto.Doctor, subsMap map[int64]dto.SubscribersInfoDTO) {
	for doctorID, subs := range subsMap {
		doc, ok := doctorsMap[doctorID]
		if !ok {
			continue
		}
		doc.InstSubsCountText = subs.InstSubsCountText
		doc.InstSubsCount = subs.InstSubsCount
		doc.TgSubsCountText = subs.TgSubsCountText
		doc.TgSubsCount = subs.TgSubsCount

		doctorsMap[doctorID] = doc
	}
}
