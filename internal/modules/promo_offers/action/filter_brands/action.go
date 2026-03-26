package filter_brands

import (
	"context"
	"medblogers_base/internal/modules/promo_offers/service/image"

	"github.com/samber/lo"

	actionDal "medblogers_base/internal/modules/promo_offers/action/filter_brands/dal"
	"medblogers_base/internal/modules/promo_offers/action/filter_brands/dto"
	"medblogers_base/internal/modules/promo_offers/client"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	commonDAO "medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	FilterBrands(ctx context.Context, req dto.BrandFilter) (brandDomain.Brands, error)
}

type CommonDal interface {
	GetBusinessCategoriesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBrandSocialNetworks(ctx context.Context, brandIDs []int64) (map[int64][]commonDAO.BrandSocialNetworkDAO, error)
}

type Action struct {
	repository   ActionDal
	commonDal    CommonDal
	brandsPhotos *image.Service
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator) *Action {
	return &Action{
		repository:   actionDal.NewRepository(pool),
		commonDal:    commonDal.NewRepository(pool),
		brandsPhotos: image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, req dto.BrandFilter) (dto.Response, error) {
	logger.Message(ctx, "[PromoOffers][FilterBrands] Фильтрация брендов")

	brands, err := a.repository.FilterBrands(ctx, req)
	if err != nil {
		return dto.Response{}, err
	}

	brandIDs := brands.IDs()
	businessCategoryIDs := brands.BusinessCategoryIDs()

	businessCategoriesMap, err := a.commonDal.GetBusinessCategoriesByIDs(ctx, businessCategoryIDs)
	if err != nil {
		return dto.Response{}, err
	}

	socialsMap, err := a.commonDal.GetBrandSocialNetworks(ctx, brandIDs)
	if err != nil {
		return dto.Response{}, err
	}

	photosMap := a.brandsPhotos.GetBrandsPhotos(ctx)

	resp := dto.Response{
		Brands: lo.Map(brands, func(item *brandDomain.Brand, _ int) dto.Brand {
			brandItem := dto.Brand{
				ID:          item.GetID(),
				Title:       item.GetTitle(),
				Slug:        item.GetSlug(),
				Photo:       photosMap[item.GetPhoto()],
				Website:     item.GetWebsite(),
				Description: item.GetDescription(),
			}

			if topicName, ok := businessCategoriesMap[item.GetBusinessCategoryID()]; ok {
				brandItem.BusinessCategory = &dto.BusinessCategory{
					ID:   item.GetBusinessCategoryID(),
					Name: topicName,
				}
			}

			brandItem.SocialNetworks = lo.Map(socialsMap[item.GetID()], func(social commonDAO.BrandSocialNetworkDAO, _ int) dto.SocialNetwork {
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
