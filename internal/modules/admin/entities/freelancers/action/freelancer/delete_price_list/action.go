package delete_price_list

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete_price_list/dal"
	commondal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
}

type ActionDal interface {
	DeletePriceList(ctx context.Context, freelancerID, priceListID int64) error
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: dal.NewRepository(pool),
		commonDal: commondal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, freelancerID, priceListID int64) error {
	_, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return err
	}

	return a.actionDal.DeletePriceList(ctx, freelancerID, priceListID)
}
