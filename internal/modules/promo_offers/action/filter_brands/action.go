package filter_brands

import (
	"context"

	"github.com/samber/lo"

	actionDal "medblogers_base/internal/modules/promo_offers/action/filter_brands/dal"
	"medblogers_base/internal/modules/promo_offers/action/filter_brands/dto"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	"medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	repository *actionDal.Repository
	commonDal  *commonDal.Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		repository: actionDal.NewRepository(pool),
		commonDal:  commonDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, req dto.BrandFilter) (dto.Response, error) {
	logger.Message(ctx, "[PromoOffers][FilterBrands] Фильтрация брендов")

	brands, err := a.repository.FilterBrands(ctx, req)
	if err != nil {
		return dto.Response{}, err
	}

	brandIDs := brands.IDs()
	topicIDs := brands.TopicIDs()

	topicsMap, err := a.commonDal.GetTopicsByIDs(ctx, topicIDs)
	if err != nil {
		return dto.Response{}, err
	}

	socialsMap, err := a.commonDal.GetBrandSocialNetworks(ctx, brandIDs)
	if err != nil {
		return dto.Response{}, err
	}

	resp := dto.Response{
		Brands: lo.Map(brands, func(item *brandDomain.Brand, _ int) dto.Brand {
			brandItem := dto.Brand{
				ID:          item.GetID(),
				Title:       item.GetTitle(),
				Slug:        item.GetSlug(),
				Photo:       item.GetPhoto(),
				Website:     item.GetWebsite(),
				Description: item.GetDescription(),
			}

			if topicName, ok := topicsMap[item.GetTopicID()]; ok {
				brandItem.Topic = &dto.Topic{
					ID:   item.GetTopicID(),
					Name: topicName,
				}
			}

			brandItem.SocialNetworks = lo.Map(socialsMap[item.GetID()], func(social dao.BrandSocialNetworkDAO, _ int) dto.SocialNetwork {
				return dto.SocialNetwork{
					ID:   social.SocialNetworkID,
					Name: social.Name,
					Slug: social.Slug,
					Link: social.Link,
				}
			})

			return brandItem
		}),
	}

	return resp, nil
}
