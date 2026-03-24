package brand_detail

import (
	"context"

	actionDal "medblogers_base/internal/modules/promo_offers/action/brand_detail/dal"
	"medblogers_base/internal/modules/promo_offers/action/brand_detail/dto"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	commonDAO "medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetBrandBySlug(ctx context.Context, slug string) (*brandDomain.Brand, error)
}

type CommonDal interface {
	GetTopicsByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBrandSocialNetworks(ctx context.Context, brandIDs []int64) (map[int64][]commonDAO.BrandSocialNetworkDAO, error)
}

type Action struct {
	repository ActionDal
	commonDal  CommonDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		repository: actionDal.NewRepository(pool),
		commonDal:  commonDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, slug string) (*dto.Brand, error) {
	logger.Message(ctx, "[PromoOffers][BrandDetail] Получение карточки бренда")

	brand, err := a.repository.GetBrandBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	topicsMap, err := a.commonDal.GetTopicsByIDs(ctx, []int64{brand.GetTopicID()})
	if err != nil {
		return nil, err
	}

	socialsMap, err := a.commonDal.GetBrandSocialNetworks(ctx, []int64{brand.GetID()})
	if err != nil {
		return nil, err
	}

	resp := &dto.Brand{
		ID:          brand.GetID(),
		Title:       brand.GetTitle(),
		Slug:        brand.GetSlug(),
		Photo:       brand.GetPhoto(),
		Website:     brand.GetWebsite(),
		Description: brand.GetDescription(),
	}

	if topicName, ok := topicsMap[brand.GetTopicID()]; ok {
		resp.Topic = &dto.Topic{
			ID:   brand.GetTopicID(),
			Name: topicName,
		}
	}

	for _, social := range socialsMap[brand.GetID()] {
		resp.SocialNetworks = append(resp.SocialNetworks, dto.SocialNetwork{
			ID:   social.SocialNetworkID,
			Name: social.Name,
			Slug: social.Slug,
			Link: social.Link,
		})
	}

	return resp, nil
}
