package image

import (
	"context"
	"io"
)

type Image interface {
	PutBrandPhoto(ctx context.Context, file io.Reader, filename string) (string, error)
	DelBrandPhoto(ctx context.Context, filename string) error
	GetBrandPhotoLink(s3Key string) string
}

type Service struct {
	image Image
}

func New(image Image) *Service {
	return &Service{image: image}
}

func (s *Service) DeleteImage(ctx context.Context, filename string) error {
	return s.image.DelBrandPhoto(ctx, filename)
}

func (s *Service) SetImage(ctx context.Context, file io.Reader, filename string) (string, error) {
	return s.image.PutBrandPhoto(ctx, file, filename)
}

func (s *Service) GetImageURL(s3Key string) string {
	return s.image.GetBrandPhotoLink(s3Key)
}
