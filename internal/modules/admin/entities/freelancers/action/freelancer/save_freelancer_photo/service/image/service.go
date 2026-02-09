package image

import (
	"context"
	"io"
)

type Image interface {
	PutFreelancerPhoto(ctx context.Context, file io.Reader, filename string) (string, error)
	DelFreelancerPhoto(ctx context.Context, filename string) error
	GetFreelancerPhotoLink(s3Key string) string
}

type Service struct {
	image Image
}

func New(image Image) *Service {
	return &Service{
		image: image,
	}
}

func (s *Service) DeleteImage(ctx context.Context, FreelancerSlug string) error {
	return s.image.DelFreelancerPhoto(ctx, FreelancerSlug)
}

func (s *Service) SetImage(ctx context.Context, file io.Reader, FreelancerSlug string) (string, error) {
	return s.image.PutFreelancerPhoto(ctx, file, FreelancerSlug)
}

func (s *Service) GetImageURL(imageLink string) string {
	return s.image.GetFreelancerPhotoLink(imageLink)
}
