package image

import (
	"context"
	"io"
)

type Image interface {
	PutSettingsPhoto(ctx context.Context, file io.Reader, filename string) (string, error)
	DelSettingsPhoto(ctx context.Context, filename string) error
	GetSettingsPhotoLink(s3Key string) string
}

type Service struct {
	image Image
}

func New(image Image) *Service {
	return &Service{
		image: image,
	}
}

func (s *Service) DeleteImage(ctx context.Context, filename string) error {
	return s.image.DelSettingsPhoto(ctx, filename)
}

func (s *Service) SetImage(ctx context.Context, file io.Reader, filename string) (string, error) {
	return s.image.PutSettingsPhoto(ctx, file, filename)
}

func (s *Service) GetImageURL(imageLink string) string {
	return s.image.GetSettingsPhotoLink(imageLink)
}
