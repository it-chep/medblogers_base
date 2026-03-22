package deactivate

import (
	"context"

	"github.com/pkg/errors"

	deactivateDal "medblogers_base/internal/modules/admin/entities/banner/action/deactivate/dal"
	commonDal "medblogers_base/internal/modules/admin/entities/banner/dal"
	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type commonGetter interface {
	GetBannerByID(ctx context.Context, bannerID int64) (*dto.Banner, error)
}

type bannerDeactivator interface {
	DeactivateBanner(ctx context.Context, bannerID int64) error
}

type Action struct {
	commonGetter      commonGetter
	bannerDeactivator bannerDeactivator
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonGetter:      commonDal.NewRepository(pool),
		bannerDeactivator: deactivateDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, bannerID int64) error {
	banner, err := a.commonGetter.GetBannerByID(ctx, bannerID)
	if err != nil {
		return err
	}

	if !banner.IsActive {
		return errors.New("Баннер уже неактивен")
	}

	return a.bannerDeactivator.DeactivateBanner(ctx, bannerID)
}
