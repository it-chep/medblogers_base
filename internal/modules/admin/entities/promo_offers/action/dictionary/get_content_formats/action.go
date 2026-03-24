package get_content_formats

import (
	"context"

	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/get_content_formats/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetContentFormats(ctx context.Context) (dictionary.NamedItems, error)
}

type Action struct {
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: actionDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) (dictionary.NamedItems, error) {
	return a.actionDal.GetContentFormats(ctx)
}
