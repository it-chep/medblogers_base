package get_by_id

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/modules/admin/client"
	getDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/get/dto"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/service/image"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	commonDAO "medblogers_base/internal/modules/admin/entities/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetOfferByID(ctx context.Context, offerID uuid.UUID) (*offerDomain.Offer, error)
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
		actionDal: commonDal.NewRepository(pool),
		commonDal: commonDal.NewRepository(pool),
		image:     image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, offerID uuid.UUID) (*getDTO.Offer, error) {
	item, err := a.actionDal.GetOfferByID(ctx, offerID)
	if err != nil {
		return nil, err
	}

	socialsMap, err := a.commonDal.GetOfferSocialNetworks(ctx, []uuid.UUID{item.GetID()})
	if err != nil {
		return nil, err
	}

	brandsMap, err := a.commonDal.GetBrandsByIDs(ctx, []int64{item.GetBrandID()})
	if err != nil {
		return nil, err
	}

	cooperationTypesMap, err := a.commonDal.GetCooperationTypesByIDs(ctx, []int64{item.GetCooperationTypeID()})
	if err != nil {
		return nil, err
	}

	businessCategoriesMap, err := a.commonDal.GetBusinessCategoriesByIDs(ctx, []int64{item.GetBusinessCategoryID()})
	if err != nil {
		return nil, err
	}

	contentFormatsMap, err := a.commonDal.GetContentFormatsByIDs(ctx, []int64{item.GetContentFormatID()})
	if err != nil {
		return nil, err
	}

	result := &getDTO.Offer{
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
		result.CooperationType = &getDTO.NamedItem{ID: item.GetCooperationTypeID(), Name: name}
	}

	if name, ok := businessCategoriesMap[item.GetBusinessCategoryID()]; ok {
		result.BusinessCategory = &getDTO.NamedItem{ID: item.GetBusinessCategoryID(), Name: name}
	}

	if name, ok := contentFormatsMap[item.GetContentFormatID()]; ok {
		result.ContentFormat = &getDTO.NamedItem{ID: item.GetContentFormatID(), Name: name}
	}

	if brandItem, ok := brandsMap[item.GetBrandID()]; ok {
		result.Brand = &getDTO.BrandPreview{
			ID:    brandItem.GetID(),
			Title: brandItem.GetTitle(),
			Slug:  brandItem.GetSlug(),
			Photo: a.image.GetImageURL(brandItem.GetPhoto()),
		}
	}

	for _, social := range socialsMap[item.GetID()] {
		result.SocialNetworks = append(result.SocialNetworks, getDTO.SocialNetwork{
			ID:   social.SocialNetworkID,
			Name: social.Name,
			Slug: social.Slug,
		})
	}

	return result, nil
}
