package filter_offers

import (
	"context"
	"medblogers_base/internal/modules/promo_offers/client"
	"medblogers_base/internal/modules/promo_offers/service/image"

	"github.com/google/uuid"

	actionDal "medblogers_base/internal/modules/promo_offers/action/filter_offers/dal"
	"medblogers_base/internal/modules/promo_offers/action/filter_offers/dto"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	commonDAO "medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	FilterOffers(ctx context.Context, req dto.OfferFilter) ([]*offerDomain.Offer, error)
}

type CommonDal interface {
	GetOfferSocialNetworks(ctx context.Context, offerIDs []uuid.UUID) (map[uuid.UUID][]commonDAO.OfferSocialNetworkDAO, error)
	GetBrandsByIDs(ctx context.Context, ids []int64) (map[int64]*brandDomain.Brand, error)
	GetCooperationTypesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBusinessCategoriesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
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

func (a *Action) Do(ctx context.Context, req dto.OfferFilter) (dto.Response, error) {
	logger.Message(ctx, "[PromoOffers][FilterOffers] Фильтрация офферов")

	offers, err := a.repository.FilterOffers(ctx, req)
	if err != nil {
		return dto.Response{}, err
	}

	ids := collectOfferIDs(offers)

	var (
		socialsMap          map[uuid.UUID][]commonDAO.OfferSocialNetworkDAO
		brandsMap           map[int64]*brandDomain.Brand
		cooperationTypesMap map[int64]string
		categoriesMap       map[int64]string
	)

	g, _ := async.WithContext(ctx)

	g.GoWithContext(func(ctx context.Context) error {
		items, err := a.commonDal.GetOfferSocialNetworks(ctx, ids.offerIDs)
		if err != nil {
			return err
		}
		socialsMap = items
		return nil
	})

	g.GoWithContext(func(ctx context.Context) error {
		items, err := a.commonDal.GetBrandsByIDs(ctx, ids.brandIDs)
		if err != nil {
			return err
		}
		brandsMap = items
		return nil
	})

	g.GoWithContext(func(ctx context.Context) error {
		items, err := a.commonDal.GetCooperationTypesByIDs(ctx, ids.cooperationTypeIDs)
		if err != nil {
			return err
		}
		cooperationTypesMap = items
		return nil
	})

	g.GoWithContext(func(ctx context.Context) error {
		items, err := a.commonDal.GetBusinessCategoriesByIDs(ctx, ids.businessCategoryIDs)
		if err != nil {
			return err
		}
		categoriesMap = items
		return nil
	})

	if err := g.Wait(); err != nil {
		return dto.Response{}, err
	}

	photosMap := a.brandsPhotos.GetBrandsPhotos(ctx)
	resp := dto.Response{
		Offers: make([]dto.Offer, 0, len(offers)),
	}

	for _, item := range offers {
		offerItem := dto.Offer{
			Description: item.GetDescription(),
			CreatedAt:   item.GetCreatedAt(),
		}

		if name, ok := cooperationTypesMap[item.GetCooperationTypeID()]; ok {
			offerItem.CooperationType = &dto.NamedItem{ID: item.GetCooperationTypeID(), Name: name}
		}

		if name, ok := categoriesMap[item.GetBusinessCategoryID()]; ok {
			offerItem.BusinessCategory = &dto.NamedItem{ID: item.GetBusinessCategoryID(), Name: name}
		}

		if brandItem, ok := brandsMap[item.GetBrandID()]; ok {
			offerItem.Photo = photosMap[brandItem.GetPhoto()]
			offerItem.Title = brandItem.GetTitle()
			offerItem.BrandDescription = brandItem.GetDescription()
		}

		for _, social := range socialsMap[item.GetID()] {
			offerItem.SocialNetworks = append(offerItem.SocialNetworks, dto.SocialNetwork{
				ID:   social.SocialNetworkID,
				Name: social.Name,
				Slug: social.Slug,
			})
		}

		resp.Offers = append(resp.Offers, offerItem)
	}

	return resp, nil
}

type offerIDs struct {
	offerIDs            []uuid.UUID
	brandIDs            []int64
	cooperationTypeIDs  []int64
	businessCategoryIDs []int64
}

func collectOfferIDs(offers []*offerDomain.Offer) offerIDs {
	result := offerIDs{
		offerIDs:            make([]uuid.UUID, 0, len(offers)),
		brandIDs:            make([]int64, 0, len(offers)),
		cooperationTypeIDs:  make([]int64, 0, len(offers)),
		businessCategoryIDs: make([]int64, 0, len(offers)),
	}

	for _, item := range offers {
		result.offerIDs = append(result.offerIDs, item.GetID())
		if item.GetBrandID() > 0 {
			result.brandIDs = append(result.brandIDs, item.GetBrandID())
		}
		if item.GetCooperationTypeID() > 0 {
			result.cooperationTypeIDs = append(result.cooperationTypeIDs, item.GetCooperationTypeID())
		}
		if item.GetBusinessCategoryID() > 0 {
			result.businessCategoryIDs = append(result.businessCategoryIDs, item.GetBusinessCategoryID())
		}
	}

	return result
}
