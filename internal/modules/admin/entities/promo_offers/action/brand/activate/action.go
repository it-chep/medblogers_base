package activate

import (
	"context"

	"github.com/pkg/errors"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/activate/dal"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetBrandByID(ctx context.Context, brandID int64) (*brandDomain.Brand, error)
}

type ActionDal interface {
	ActivateBrand(ctx context.Context, brandID int64) error
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: commonDal.NewRepository(pool),
		actionDal: actionDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, brandID int64) error {
	item, err := a.commonDal.GetBrandByID(ctx, brandID)
	if err != nil {
		return err
	}

	if item.GetIsActive() {
		return errors.New("Бренд уже активен")
	}

	return a.actionDal.ActivateBrand(ctx, brandID)
}
