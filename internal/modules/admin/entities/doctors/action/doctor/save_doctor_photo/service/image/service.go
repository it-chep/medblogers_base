package image

import (
	"context"
	"io"
)

type Image interface {
	PutDoctorPhoto(ctx context.Context, file io.Reader, filename string) (string, error)
	DelDoctorPhoto(ctx context.Context, filename string) error
}

type Service struct {
	image Image
}

func New(image Image) *Service {
	return &Service{
		image: image,
	}
}

func (s *Service) DeleteImage(ctx context.Context, doctorSlug string) error {
	return s.image.DelDoctorPhoto(ctx, doctorSlug)
}

func (s *Service) SetImage(ctx context.Context, file io.Reader, doctorSlug string) (string, error) {
	return s.image.PutDoctorPhoto(ctx, file, doctorSlug)
}
