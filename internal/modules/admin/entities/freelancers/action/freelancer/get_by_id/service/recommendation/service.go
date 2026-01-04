package recommendation

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/doctor"
)

type Dal interface {
	GetRecommendations(ctx context.Context, freelancerID int64) ([]int64, error)
	GetRecommendationInfoByIDs(ctx context.Context, doctorIDs []int64) ([]*doctor.Doctor, error)
}

type Service struct {
	dal Dal
}

func New(dal Dal) *Service {
	return &Service{
		dal: dal,
	}
}
func (s *Service) Enrich(ctx context.Context, freelancerDTO *dto.FreelancerDTO) (*dto.FreelancerDTO, error) {
	doctorIDs, err := s.dal.GetRecommendations(ctx, freelancerDTO.ID)
	if err != nil {
		return freelancerDTO, err
	}

	recommendations, err := s.dal.GetRecommendationInfoByIDs(ctx, doctorIDs)
	if err != nil {
		return freelancerDTO, err
	}

	freelancerDTO.Recommendations = lo.Map(recommendations, func(item *doctor.Doctor, _ int) dto.Recommendation {
		return dto.Recommendation{
			DoctorID:   int64(item.GetID()),
			DoctorName: item.GetName(),
		}
	})

	return freelancerDTO, nil
}
