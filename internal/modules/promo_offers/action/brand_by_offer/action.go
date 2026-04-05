package brand_by_offer

import (
	"context"

	actionDal "medblogers_base/internal/modules/promo_offers/action/brand_by_offer/dal"
	"medblogers_base/internal/modules/promo_offers/action/brand_by_offer/dto"
	"medblogers_base/internal/modules/promo_offers/client"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	commonDAO "medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	"medblogers_base/internal/modules/promo_offers/service/image"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetBrandByOffer(ctx context.Context, offerID string) (*brandDomain.Brand, error)
}

type CommonDal interface {
	GetBusinessCategoriesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBrandSocialNetworks(ctx context.Context, brandIDs []int64) (map[int64][]commonDAO.BrandSocialNetworkDAO, error)
}

type Action struct {
	repository ActionDal
	commonDal  CommonDal
	image      *image.Service
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator) *Action {
	return &Action{
		repository: actionDal.NewRepository(pool),
		commonDal:  commonDal.NewRepository(pool),
		image:      image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, offerID string) (*dto.Brand, error) {
	logger.Message(ctx, "[PromoOffers][BrandByOffer] Получение карточки бренда по ОФФЕРУ")

	brand, err := a.repository.GetBrandByOffer(ctx, offerID)
	if err != nil {
		return nil, err
	}

	businessCategoriesMap, err := a.commonDal.GetBusinessCategoriesByIDs(ctx, []int64{brand.GetBusinessCategoryID()})
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
		Photo:       a.image.EnrichPhotoByKey(brand.GetPhoto()),
		Website:     brand.GetWebsite(),
		Description: brand.GetDescription(),
	}

	if topicName, ok := businessCategoriesMap[brand.GetBusinessCategoryID()]; ok {
		resp.BusinessCategory = &dto.BusinessCategory{
			ID:   brand.GetBusinessCategoryID(),
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
