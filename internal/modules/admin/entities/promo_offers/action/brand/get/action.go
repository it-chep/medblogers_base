package get

import (
	"context"

	"medblogers_base/internal/modules/admin/client"
	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/get/dto"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/service/image"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetBrands(ctx context.Context) (brandDomain.Brands, error)
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
		actionDal: actionDal.NewRepository(pool),
		commonDal: commonDal.NewRepository(pool),
		image:     image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context) ([]dto.Brand, error) {
	brands, err := a.actionDal.GetBrands(ctx)
	if err != nil {
		return nil, err
	}

	topicsMap, err := a.commonDal.GetTopicsByIDs(ctx, brands.TopicIDs())
	if err != nil {
		return nil, err
	}

	socialsMap, err := a.commonDal.GetBrandSocialNetworks(ctx, brands.IDs())
	if err != nil {
		return nil, err
	}

	result := make([]dto.Brand, 0, len(brands))
	for _, item := range brands {
		result = append(result, dto.NewBrand(
			item,
			topicsMap[item.GetTopicID()],
			socialsMap[item.GetID()],
			a.image.GetImageURL(item.GetPhoto()),
		))
	}

	return result, nil
}
