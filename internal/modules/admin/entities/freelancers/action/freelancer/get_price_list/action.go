package get_price_list

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_price_list/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_price_list/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Dal interface {
	GetPriceList(ctx context.Context, freelancerID int64) ([]dto.PriceList, error)
}

type Action struct {
	dal Dal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do получение прайслиста конкретного фрилансера
func (a *Action) Do(ctx context.Context, freelancerID int64) ([]dto.PriceList, error) {
	return a.dal.GetPriceList(ctx, freelancerID)
}
