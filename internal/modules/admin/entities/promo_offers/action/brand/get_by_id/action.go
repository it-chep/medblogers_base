package get_by_id

import (
	"context"

	"medblogers_base/internal/modules/admin/client"
	getDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get/dto"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/service/image"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetBrandByID(ctx context.Context, brandID int64) (*brandDomain.Brand, error)
}

type CommonDal interface {
	GetTopicsByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBrandSocialNetworks(ctx context.Context, brandIDs []int64) (map[int64]dictionary.BrandSocialNetworks, error)
}

type Action struct {
	actionDal ActionDal
	commonDal CommonDal
	image     *image.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal: commonDal.NewRepository(pool),
		commonDal: commonDal.NewRepository(pool),
		image:     image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, brandID int64) (*getDTO.Brand, error) {
	item, err := a.actionDal.GetBrandByID(ctx, brandID)
	if err != nil {
		return nil, err
	}

	topicsMap, err := a.commonDal.GetTopicsByIDs(ctx, []int64{item.GetTopicID()})
	if err != nil {
		return nil, err
	}

	socialsMap, err := a.commonDal.GetBrandSocialNetworks(ctx, []int64{item.GetID()})
	if err != nil {
		return nil, err
	}

	result := getDTO.NewBrand(
		item,
		topicsMap[item.GetTopicID()],
		socialsMap[item.GetID()],
		a.image.GetImageURL(item.GetPhoto()),
	)

	return &result, nil
}
