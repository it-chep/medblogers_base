package image

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
)

type ImageGetter interface {
	GetDoctorPhotoLink(s3Key string) string
}

type Service struct {
	imageGetter ImageGetter
}

func New(imageGetter ImageGetter) *Service {
	return &Service{
		imageGetter: imageGetter,
	}
}

func (s *Service) Enrich(_ context.Context, docDTO *dto.DoctorDTO) (*dto.DoctorDTO, error) {
	if len(docDTO.S3Key.String()) == 0 {
		return docDTO, nil
	}
	doctorImageURL := s.imageGetter.GetDoctorPhotoLink(docDTO.S3Key.String())

	docDTO.Image = doctorImageURL

	return docDTO, nil
}
