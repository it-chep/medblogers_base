package activate

import (
	"context"

	"github.com/pkg/errors"

	activateDal "medblogers_base/internal/modules/admin/entities/banner/action/activate/dal"
	commonDal "medblogers_base/internal/modules/admin/entities/banner/dal"
	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type commonGetter interface {
	GetBannerByID(ctx context.Context, bannerID int64) (*dto.Banner, error)
}

type bannerActivator interface {
	ActivateBanner(ctx context.Context, bannerID int64) error
}

type Action struct {
	commonGetter    commonGetter
	bannerActivator bannerActivator
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonGetter:    commonDal.NewRepository(pool),
		bannerActivator: activateDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, bannerID int64) error {
	banner, err := a.commonGetter.GetBannerByID(ctx, bannerID)
	if err != nil {
		return err
	}

	if banner.IsActive {
		return errors.New("Баннер уже активен")
	}

	return a.bannerActivator.ActivateBanner(ctx, bannerID)
}
