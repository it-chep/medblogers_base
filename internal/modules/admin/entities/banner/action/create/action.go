package create

import (
	"context"

	createDal "medblogers_base/internal/modules/admin/entities/banner/action/create/dal"
	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type bannerCreator interface {
	CreateBanner(ctx context.Context, req dto.UpdateRequest) (int64, error)
}

type Action struct {
	bannerCreator bannerCreator
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		bannerCreator: createDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, req dto.UpdateRequest) (int64, error) {
	return a.bannerCreator.CreateBanner(ctx, req)
}
