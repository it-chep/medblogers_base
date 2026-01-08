package add_price_list

import (
	"context"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_price_list/dal"
	commondal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
}

type ActionDal interface {
	AddPriceList(ctx context.Context, freelancerID int64, name string, price int64) error
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

func (a *Action) Do(ctx context.Context, freelancerID int64, name string, price int64) error {
	_, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return err
	}

	if price < 0 {
		return errors.New("invalid price")
	}
	return a.actionDal.AddPriceList(ctx, freelancerID, name, price)
}
