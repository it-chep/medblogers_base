package subscribers

import (
	"context"
	"medblogers_base/internal/modules/admin/client/subscribers/indto"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
)

type SubscribersGetter interface {
	GetDoctorSubscribers(ctx context.Context, medblogersID int64) (indto.GetDoctorSubscribersResponse, error)
}

type Service struct {
	getter SubscribersGetter
}

func New(getter SubscribersGetter) *Service {
	return &Service{
		getter: getter,
	}
}

func (s *Service) Enrich(ctx context.Context, docDTO *dto.DoctorDTO) (*dto.DoctorDTO, error) {

	subscribers, err := s.getter.GetDoctorSubscribers(ctx, docDTO.ID)
	if err != nil {
		return nil, err
	}

	docDTO.SubscribersInfo = []dto.Subscribers{}
	if len(docDTO.InstURL) != 0 {
		docDTO.SubscribersInfo = append(docDTO.SubscribersInfo, dto.Subscribers{
			Key:             "inst",
			SubsCount:       subscribers.InstSubsCount,
			SubsCountText:   subscribers.InstSubsCountText,
			LastUpdatedDate: subscribers.InstLastUpdatedDate,
		})
	}

	if len(docDTO.YoutubeURL) != 0 {
		docDTO.SubscribersInfo = append(docDTO.SubscribersInfo, dto.Subscribers{
			Key:             "youtube",
			SubsCount:       subscribers.YouTubeSubsCount,
			SubsCountText:   subscribers.YouTubeSubsCountText,
			LastUpdatedDate: subscribers.YouTubeLastUpdatedDate,
		})
	}

	if len(docDTO.VkURL) != 0 {
		docDTO.SubscribersInfo = append(docDTO.SubscribersInfo, dto.Subscribers{
			Key:             "vk",
			SubsCount:       subscribers.VkSubsCount,
			SubsCountText:   subscribers.VkSubsCountText,
			LastUpdatedDate: subscribers.VkLastUpdatedDate,
		})
	}

	if len(docDTO.TgChannelURL) != 0 {
		docDTO.SubscribersInfo = append(docDTO.SubscribersInfo, dto.Subscribers{
			Key:             "tg",
			SubsCount:       subscribers.TgSubsCount,
			SubsCountText:   subscribers.TgSubsCountText,
			LastUpdatedDate: subscribers.TgLastUpdatedDate,
		})
	}

	return docDTO, nil
}
