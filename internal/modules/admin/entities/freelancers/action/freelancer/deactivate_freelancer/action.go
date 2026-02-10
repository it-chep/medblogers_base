package deactivate_freelancer

import (
	"context"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/deactivate_freelancer/dal"
	common_dal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
}

type ActionDal interface {
	DeactivateFreelancer(ctx context.Context, freelancerID int64) (err error)
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

func (a *Action) Do(ctx context.Context, freelancerID int64) error {
	frncer, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return err
	}

	if !frncer.GetIsActive() {
		return errors.New("Фрилансер уже деактивирован")
	}

	return a.actionDal.DeactivateFreelancer(ctx, freelancerID)
}
