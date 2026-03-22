package update

import (
	"context"

	updateDal "medblogers_base/internal/modules/admin/entities/banner/action/update/dal"
	commonDal "medblogers_base/internal/modules/admin/entities/banner/dal"
	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type commonGetter interface {
	GetBannerByID(ctx context.Context, bannerID int64) (*dto.Banner, error)
}

type bannerUpdater interface {
	UpdateBanner(ctx context.Context, bannerID int64, req dto.UpdateRequest) error
}

type Action struct {
	commonGetter  commonGetter
	bannerUpdater bannerUpdater
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonGetter:  commonDal.NewRepository(pool),
		bannerUpdater: updateDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, bannerID int64, req dto.UpdateRequest) error {
	if _, err := a.commonGetter.GetBannerByID(ctx, bannerID); err != nil {
		return err
	}

	return a.bannerUpdater.UpdateBanner(ctx, bannerID, req)
}
