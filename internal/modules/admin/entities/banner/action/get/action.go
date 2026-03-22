package get

import (
	"context"

	commonDal "medblogers_base/internal/modules/admin/entities/banner/dal"
	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type bannersGetter interface {
	GetBanners(ctx context.Context) ([]dto.Banner, error)
}

type Action struct {
	bannersGetter bannersGetter
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		bannersGetter: commonDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]dto.Banner, error) {
	return a.bannersGetter.GetBanners(ctx)
}
