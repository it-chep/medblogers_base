package get_by_id

import (
	"context"

	commonDal "medblogers_base/internal/modules/admin/entities/banner/dal"
	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type bannerGetter interface {
	GetBannerByID(ctx context.Context, bannerID int64) (*dto.Banner, error)
}

type Action struct {
	bannerGetter bannerGetter
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		bannerGetter: commonDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, bannerID int64) (*dto.Banner, error) {
	return a.bannerGetter.GetBannerByID(ctx, bannerID)
}
