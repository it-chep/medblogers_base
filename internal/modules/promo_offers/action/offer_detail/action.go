package offer_detail

import (
	"context"
	"medblogers_base/internal/modules/promo_offers/client"
	"medblogers_base/internal/modules/promo_offers/service/image"

	"github.com/google/uuid"

	actionDal "medblogers_base/internal/modules/promo_offers/action/offer_detail/dal"
	"medblogers_base/internal/modules/promo_offers/action/offer_detail/dto"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	commonDAO "medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetOfferByID(ctx context.Context, id uuid.UUID) (*offerDomain.Offer, error)
}

type CommonDal interface {
	GetCooperationTypesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBusinessCategoriesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBrandsByIDs(ctx context.Context, ids []int64) (map[int64]*brandDomain.Brand, error)
	GetBrandSocialNetworks(ctx context.Context, brandIDs []int64) (map[int64][]commonDAO.BrandSocialNetworkDAO, error)
	GetOfferSocialNetworks(ctx context.Context, offerIDs []uuid.UUID) (map[uuid.UUID][]commonDAO.OfferSocialNetworkDAO, error)
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

func (a *Action) Do(ctx context.Context, id uuid.UUID) (*dto.Offer, error) {
	var (
		cooperationTypesMap map[int64]string
		brandsMap           map[int64]*brandDomain.Brand
		brandSocialsMap     map[int64][]commonDAO.BrandSocialNetworkDAO
		socialsMap          map[uuid.UUID][]commonDAO.OfferSocialNetworkDAO
	)

	logger.Message(ctx, "[PromoOffers][OfferDetail] Получение карточки оффера")

	offer, err := a.repository.GetOfferByID(ctx, id)
	if err != nil {
		return nil, err
	}

	g := async.NewGroup()

	g.Go(func() {
		items, gErr := a.commonDal.GetCooperationTypesByIDs(ctx, []int64{offer.GetCooperationTypeID()})
		if gErr != nil {
			logger.Error(ctx, "[PromoOffers][OfferDetail] Ошибка при получении типа сотрудничества", gErr)
			return
		}

		cooperationTypesMap = items
	})

	g.Go(func() {
		items, gErr := a.commonDal.GetBrandsByIDs(ctx, []int64{offer.GetBrandID()})
		if gErr != nil {
			logger.Error(ctx, "[PromoOffers][OfferDetail] Ошибка при получении бренда", gErr)
			return
		}

		brandsMap = items
	})

	g.Go(func() {
		items, gErr := a.commonDal.GetBrandSocialNetworks(ctx, []int64{offer.GetBrandID()})
		if gErr != nil {
			logger.Error(ctx, "[PromoOffers][OfferDetail] Ошибка при получении соцсетей бренда", gErr)
			return
		}

		brandSocialsMap = items
	})

	g.Go(func() {
		items, gErr := a.commonDal.GetOfferSocialNetworks(ctx, []uuid.UUID{offer.GetID()})
		if gErr != nil {
			logger.Error(ctx, "[PromoOffers][OfferDetail] Ошибка при получении соцсетей оффера", gErr)
			return
		}

		socialsMap = items
	})

	g.Wait()

	resp := &dto.Offer{
		Description: offer.GetDescription(),
		Price:       offer.GetPrice(),
		CreatedAt:   offer.GetCreatedAt(),
	}

	if name, ok := cooperationTypesMap[offer.GetCooperationTypeID()]; ok {
		resp.CooperationType = &dto.NamedItem{ID: offer.GetCooperationTypeID(), Name: name}
	}

	if brandItem, ok := brandsMap[offer.GetBrandID()]; ok {
		brand := &dto.Brand{
			Photo:       a.image.EnrichPhotoByKey(brandItem.GetPhoto()),
			Title:       brandItem.GetTitle(),
			Description: brandItem.GetDescription(),
			About:       brandItem.GetAbout(),
		}

		if brandItem.GetBusinessCategoryID() > 0 {
			businessCategoriesMap, err := a.commonDal.GetBusinessCategoriesByIDs(ctx, []int64{brandItem.GetBusinessCategoryID()})
			if err != nil {
				logger.Error(ctx, "[PromoOffers][OfferDetail] Ошибка при получении темы бренда", err)
			} else if name, ok := businessCategoriesMap[brandItem.GetBusinessCategoryID()]; ok {
				brand.BusinessCategory = &dto.NamedItem{ID: brandItem.GetBusinessCategoryID(), Name: name}
			}
		}

		for _, social := range brandSocialsMap[brandItem.GetID()] {
			brand.SocialNetworks = append(brand.SocialNetworks, dto.BrandSocialNetwork{
				ID:   social.SocialNetworkID,
				Name: social.Name,
				Slug: social.Slug,
				Link: social.Link,
			})
		}

		resp.Brand = brand
	}

	for _, social := range socialsMap[offer.GetID()] {
		resp.SocialNetworks = append(resp.SocialNetworks, dto.SocialNetwork{
			ID:   social.SocialNetworkID,
			Name: social.Name,
			Slug: social.Slug,
		})
	}

	return resp, nil
}
