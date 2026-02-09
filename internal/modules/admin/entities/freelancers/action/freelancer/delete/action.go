package delete

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	DeleteFreelancer(ctx context.Context, id int64) error
	DeleteFreelancerRecommendation(ctx context.Context, id int64) error
	DeleteFreelancerNetworks(ctx context.Context, id int64) error
	DeleteFreelancerCities(ctx context.Context, id int64) error
	DeleteFreelancerPriceList(ctx context.Context, id int64) error
	DeleteFreelancerSpecialities(ctx context.Context, id int64) error
}

type Action struct {
	actionDal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID int64) (err error) {
	// todo траназакицю ?
	err = a.actionDal.DeleteFreelancerNetworks(ctx, freelancerID)
	if err != nil {
		return err
	}
	err = a.actionDal.DeleteFreelancerRecommendation(ctx, freelancerID)
	if err != nil {
		return err
	}
	err = a.actionDal.DeleteFreelancerSpecialities(ctx, freelancerID)
	if err != nil {
		return err
	}
	err = a.actionDal.DeleteFreelancerCities(ctx, freelancerID)
	if err != nil {
		return err
	}
	err = a.actionDal.DeleteFreelancerPriceList(ctx, freelancerID)
	if err != nil {
		return err
	}
	return a.actionDal.DeleteFreelancer(ctx, freelancerID)
}
