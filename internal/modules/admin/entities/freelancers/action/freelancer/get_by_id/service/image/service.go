package image

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
)

type Getter interface {
	GetFreelancerPhotoLink(s3Key string) string
}

type Service struct {
	getter Getter
}

func New(getter Getter) *Service {
	return &Service{
		getter: getter,
	}
}

func (s *Service) Enrich(_ context.Context, freelancerDTO *dto.FreelancerDTO) (*dto.FreelancerDTO, error) {
	freelancerDTO.S3Image = s.getter.GetFreelancerPhotoLink(freelancerDTO.S3Key)
	return freelancerDTO, nil
}
