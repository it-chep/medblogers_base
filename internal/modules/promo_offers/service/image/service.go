package image

import (
	"context"
	"medblogers_base/internal/pkg/logger"
)

type ImageGetter interface {
	GetBrandPhotoLink(s3Key string) string
	GetBrandsPhotos(ctx context.Context) (map[string]string, error)
}

type Service struct {
	imageGetter ImageGetter
}

func New(imageGetter ImageGetter) *Service {
	return &Service{imageGetter: imageGetter}
}

// EnrichPhotoByKey получение фотографии конкретного бренда
func (s *Service) EnrichPhotoByKey(s3Key string) string {
	if s3Key == "" {
		return ""
	}

	return s.imageGetter.GetBrandPhotoLink(s3Key)
}

// GetBrandsPhotos получение фотографий брендов
func (s *Service) GetBrandsPhotos(ctx context.Context) map[string]string {
	photos, err := s.imageGetter.GetBrandsPhotos(ctx)
	if err != nil {
		logger.Error(ctx, "ошибка при получении фотографий брендов", err)
		return nil
	}

	return photos
}
