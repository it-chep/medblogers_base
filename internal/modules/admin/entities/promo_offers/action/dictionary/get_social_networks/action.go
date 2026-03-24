package get_social_networks

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/get_social_networks/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetSocialNetworks(ctx context.Context) (dictionary.SocialNetworks, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: actionDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) (dictionary.SocialNetworks, error) {
	return a.actionDal.GetSocialNetworks(ctx)
}
