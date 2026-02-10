package freelancer

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/update/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
)

type CommonDal interface {
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
}

type Dal interface {
	UpdateFreelancer(ctx context.Context, freelancerID int64, req dto.UpdateRequest) error
}

type Service struct {
	commonDal CommonDal
	dal       Dal
}

func NewService(commonDal CommonDal, dal Dal) *Service {
	return &Service{
		commonDal: commonDal,
		dal:       dal,
	}
}

// GetFreelancer получение существующего фрилансера
func (s *Service) GetFreelancer(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error) {
	return s.commonDal.GetFreelancerByID(ctx, freelancerID)
}

// UpdateFreelancer обновление фрилансера
func (s *Service) UpdateFreelancer(ctx context.Context, freelancerID int64, updateReq dto.UpdateRequest) error {
	return s.dal.UpdateFreelancer(ctx, freelancerID, updateReq)
}
