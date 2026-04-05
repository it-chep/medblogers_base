package create_topic

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/create_topic/dal"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	CreateTopic(ctx context.Context, name string) (int64, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{actionDal: actionDal.NewRepository(pool)}
}

func (a *Action) Do(ctx context.Context, name string) (int64, error) {
	return a.actionDal.CreateTopic(ctx, name)
}
