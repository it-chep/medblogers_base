package filter

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/modules/admin/client"
	actionDal "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/filter/dal"
	filterDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/filter/dto"
	getDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get/dto"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/service/image"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	commonDAO "medblogers_base/internal/modules/admin/entities/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	FilterOffers(ctx context.Context, req filterDTO.Request) (offerDomain.Offers, error)
}

type CommonDal interface {
	GetOfferSocialNetworks(ctx context.Context, offerIDs []uuid.UUID) (map[uuid.UUID][]commonDAO.OfferSocialNetworkDAO, error)
	GetBrandsByIDs(ctx context.Context, ids []int64) (map[int64]*brandDomain.Brand, error)
	GetCooperationTypesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBusinessCategoriesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetContentFormatsByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
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

func (a *Action) Do(ctx context.Context, req filterDTO.Request) ([]getDTO.Offer, error) {
	offers, err := a.actionDal.FilterOffers(ctx, req)
	if err != nil {
		return nil, err
	}

	return a.enrichOffers(ctx, offers)
}

func (a *Action) enrichOffers(ctx context.Context, offers offerDomain.Offers) ([]getDTO.Offer, error) {
	offerIDs := offers.IDs()
	brandIDs := offers.BrandIDs()
	cooperationTypeIDs := offers.CooperationTypeIDs()
	businessCategoryIDs := offers.BusinessCategoryIDs()
	contentFormatIDs := offers.ContentFormatIDs()

	socialsMap, err := a.commonDal.GetOfferSocialNetworks(ctx, offerIDs)
	if err != nil {
		return nil, err
	}

	brandsMap, err := a.commonDal.GetBrandsByIDs(ctx, brandIDs)
	if err != nil {
		return nil, err
	}

	cooperationTypesMap, err := a.commonDal.GetCooperationTypesByIDs(ctx, cooperationTypeIDs)
	if err != nil {
		return nil, err
	}

	businessCategoriesMap, err := a.commonDal.GetBusinessCategoriesByIDs(ctx, businessCategoryIDs)
	if err != nil {
		return nil, err
	}

	contentFormatsMap, err := a.commonDal.GetContentFormatsByIDs(ctx, contentFormatIDs)
	if err != nil {
		return nil, err
	}

	result := make([]getDTO.Offer, 0, len(offers))
	for _, item := range offers {
		offerItem := getDTO.Offer{
			ID:                   item.GetID().String(),
			Title:                item.GetTitle(),
			Description:          item.GetDescription(),
			Price:                item.GetPrice(),
			PublicationDate:      item.GetPublicationDate(),
			AdMarkingResponsible: item.GetAdMarkingResponsible(),
			ResponsesCapacity:    item.GetResponsesCapacity(),
			IsActive:             item.GetIsActive(),
			CreatedAt:            item.GetCreatedAt(),
		}

		if name, ok := cooperationTypesMap[item.GetCooperationTypeID()]; ok {
			offerItem.CooperationType = &getDTO.NamedItem{ID: item.GetCooperationTypeID(), Name: name}
		}

		if name, ok := businessCategoriesMap[item.GetBusinessCategoryID()]; ok {
			offerItem.BusinessCategory = &getDTO.NamedItem{ID: item.GetBusinessCategoryID(), Name: name}
		}

		if name, ok := contentFormatsMap[item.GetContentFormatID()]; ok {
			offerItem.ContentFormat = &getDTO.NamedItem{ID: item.GetContentFormatID(), Name: name}
		}

		if brandItem, ok := brandsMap[item.GetBrandID()]; ok {
			offerItem.Brand = &getDTO.BrandPreview{
				ID:    brandItem.GetID(),
				Title: brandItem.GetTitle(),
				Slug:  brandItem.GetSlug(),
				Photo: a.image.GetImageURL(brandItem.GetPhoto()),
			}
		}

		for _, social := range socialsMap[item.GetID()] {
			offerItem.SocialNetworks = append(offerItem.SocialNetworks, getDTO.SocialNetwork{
				ID:   social.SocialNetworkID,
				Name: social.Name,
				Slug: social.Slug,
			})
		}

		result = append(result, offerItem)
	}

	return result, nil
}
