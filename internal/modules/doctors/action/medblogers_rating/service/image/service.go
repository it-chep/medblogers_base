package image

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/medblogers_rating/dto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/logger"
)

type ImageGetter interface {
	GetUserPhotos(ctx context.Context) (map[doctor.S3Key]string, error)
}

type Service struct {
	imageGetter ImageGetter
}

func New(imageGetter ImageGetter) *Service {
	return &Service{
		imageGetter: imageGetter,
	}
}

func (s *Service) EnrichImages(ctx context.Context, items []dto.RatingItem) {
	photos, err := s.imageGetter.GetUserPhotos(ctx)
	if err != nil {
		logger.Error(ctx, "[MedblogersRating] Ошибка при получении фотографий", err)
		for i := range items {
			items[i].Image = "https://storage.yandexcloud.net/medblogers-photos/images/zag.png"
		}
		return
	}

	for i := range items {
		photo, ok := photos[doctor.S3Key(items[i].S3Image)]
		if !ok || items[i].S3Image == "" {
			items[i].Image = "https://storage.yandexcloud.net/medblogers-photos/images/zag.png"
			continue
		}
		items[i].Image = photo
	}
}
