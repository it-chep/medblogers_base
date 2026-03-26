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
	GetTopicsByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetContentFormatsByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetBrandsByIDs(ctx context.Context, ids []int64) (map[int64]*brandDomain.Brand, error)
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
		topicsMap           map[int64]string
		contentFormatsMap   map[int64]string
		brandsMap           map[int64]*dto.BrandPreview
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
		items, gErr := a.commonDal.GetTopicsByIDs(ctx, []int64{offer.GetTopicID()})
		if gErr != nil {
			logger.Error(ctx, "[PromoOffers][OfferDetail] Ошибка при получении темы", gErr)
			return
		}

		topicsMap = items
	})

	g.Go(func() {
		items, gErr := a.commonDal.GetContentFormatsByIDs(ctx, []int64{offer.GetContentFormatID()})
		if gErr != nil {
			logger.Error(ctx, "[PromoOffers][OfferDetail] Ошибка при получении формата контента", gErr)
			return
		}

		contentFormatsMap = items
	})

	g.Go(func() {
		items, gErr := a.commonDal.GetBrandsByIDs(ctx, []int64{offer.GetBrandID()})
		if gErr != nil {
			logger.Error(ctx, "[PromoOffers][OfferDetail] Ошибка при получении бренда", gErr)
			return
		}

		mapped := make(map[int64]*dto.BrandPreview, len(items))
		for id, item := range items {
			mapped[id] = &dto.BrandPreview{
				ID:    item.GetID(),
				Title: item.GetTitle(),
				Slug:  item.GetSlug(),
				Photo: a.image.EnrichPhotoByKey(item.GetPhoto()),
			}
		}

		brandsMap = mapped
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
		ID:                   offer.GetID().String(),
		Title:                offer.GetTitle(),
		Description:          offer.GetDescription(),
		Price:                offer.GetPrice(),
		PublicationDate:      offer.GetPublicationDate(),
		AdMarkingResponsible: offer.GetAdMarkingResponsible(),
		ResponsesCapacity:    offer.GetResponsesCapacity(),
	}

	if name, ok := cooperationTypesMap[offer.GetCooperationTypeID()]; ok {
		resp.CooperationType = &dto.NamedItem{ID: offer.GetCooperationTypeID(), Name: name}
	}

	if name, ok := topicsMap[offer.GetTopicID()]; ok {
		resp.Topic = &dto.NamedItem{ID: offer.GetTopicID(), Name: name}
	}

	if name, ok := contentFormatsMap[offer.GetContentFormatID()]; ok {
		resp.ContentFormat = &dto.NamedItem{ID: offer.GetContentFormatID(), Name: name}
	}

	if brandItem, ok := brandsMap[offer.GetBrandID()]; ok {
		resp.Brand = brandItem
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
