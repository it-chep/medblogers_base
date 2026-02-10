package add_recommendation

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_recommendation/dal"
	common_dal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/doctor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetDoctorToRecommendation(ctx context.Context, doctorID int64) (*doctor.Doctor, error)
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
}

type ActionDal interface {
	AddRecommendation(ctx context.Context, freelancerId, doctorId int64) error
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
		commonDal: common_dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID, doctorID int64) error {
	_, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return err
	}

	_, err = a.commonDal.GetDoctorToRecommendation(ctx, doctorID)
	if err != nil {
		return err
	}

	return a.actionDal.AddRecommendation(ctx, freelancerID, doctorID)
}
