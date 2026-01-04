package price_list

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id/dto"
)

type Dal interface {
	GetPriceList(ctx context.Context, freelancerID int64) ([]dto.PriceList, error)
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
	priceList, err := s.dal.GetPriceList(ctx, freelancerDTO.ID)
	if err != nil {
		return freelancerDTO, err
	}

	freelancerDTO.PriceList = priceList

	return freelancerDTO, nil
}
