package image

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
)

type ImageGetter interface {
	GeneratePresignedURL(ctx context.Context, s3Key string) (string, error)
}

type Service struct {
	imageGetter ImageGetter
}

func New(imageGetter ImageGetter) *Service {
	return &Service{
		imageGetter: imageGetter,
	}
}

func (s *Service) Enrich(ctx context.Context, docDTO *dto.DoctorDTO) (*dto.DoctorDTO, error) {
	doctorImageURL, err := s.imageGetter.GeneratePresignedURL(ctx, docDTO.S3Key.String())
	if err != nil {
		return docDTO, err
	}

	docDTO.Image = doctorImageURL

	return docDTO, nil
}
