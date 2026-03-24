package brand_offers

import (
	"context"

	"github.com/google/uuid"

	actionDal "medblogers_base/internal/modules/promo_offers/action/brand_offers/dal"
	"medblogers_base/internal/modules/promo_offers/action/brand_offers/dto"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	commonDAO "medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetBrandBySlug(ctx context.Context, slug string) (*brandDomain.Brand, error)
	GetOffersByBrandID(ctx context.Context, brandID int64) ([]*offerDomain.Offer, error)
}

type CommonDal interface {
	GetOfferSocialNetworks(ctx context.Context, offerIDs []uuid.UUID) (map[uuid.UUID][]commonDAO.OfferSocialNetworkDAO, error)
	GetCooperationTypesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetTopicsByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
	GetContentFormatsByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
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

func (a *Action) Do(ctx context.Context, brandSlug string) (dto.Response, error) {
	logger.Message(ctx, "[PromoOffers][BrandOffers] Получение офферов бренда")

	brand, err := a.repository.GetBrandBySlug(ctx, brandSlug)
	if err != nil {
		return dto.Response{}, err
	}

	offers, err := a.repository.GetOffersByBrandID(ctx, brand.GetID())
	if err != nil {
		return dto.Response{}, err
	}

	ids := collectOfferIDs(offers)

	var (
		socialsMap          map[uuid.UUID][]commonDAO.OfferSocialNetworkDAO
		cooperationTypesMap map[int64]string
		topicsMap           map[int64]string
		contentFormatsMap   map[int64]string
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
		items, err := a.commonDal.GetCooperationTypesByIDs(ctx, ids.cooperationTypeIDs)
		if err != nil {
			return err
		}
		cooperationTypesMap = items
		return nil
	})

	g.GoWithContext(func(ctx context.Context) error {
		items, err := a.commonDal.GetTopicsByIDs(ctx, ids.topicIDs)
		if err != nil {
			return err
		}
		topicsMap = items
		return nil
	})

	g.GoWithContext(func(ctx context.Context) error {
		items, err := a.commonDal.GetContentFormatsByIDs(ctx, ids.contentFormatIDs)
		if err != nil {
			return err
		}
		contentFormatsMap = items
		return nil
	})

	if err := g.Wait(); err != nil {
		return dto.Response{}, err
	}

	brandPreview := &dto.BrandPreview{
		ID:    brand.GetID(),
		Title: brand.GetTitle(),
		Slug:  brand.GetSlug(),
		Photo: brand.GetPhoto(),
	}

	resp := dto.Response{
		Offers: make([]dto.Offer, 0, len(offers)),
	}

	for _, item := range offers {
		offerItem := dto.Offer{
			ID:                   item.GetID().String(),
			Title:                item.GetTitle(),
			Description:          item.GetDescription(),
			Price:                item.GetPrice(),
			PublicationDate:      item.GetPublicationDate(),
			AdMarkingResponsible: item.GetAdMarkingResponsible(),
			ResponsesCapacity:    item.GetResponsesCapacity(),
			Brand:                brandPreview,
		}

		if name, ok := cooperationTypesMap[item.GetCooperationTypeID()]; ok {
			offerItem.CooperationType = &dto.NamedItem{ID: item.GetCooperationTypeID(), Name: name}
		}

		if name, ok := topicsMap[item.GetTopicID()]; ok {
			offerItem.Topic = &dto.NamedItem{ID: item.GetTopicID(), Name: name}
		}

		if name, ok := contentFormatsMap[item.GetContentFormatID()]; ok {
			offerItem.ContentFormat = &dto.NamedItem{ID: item.GetContentFormatID(), Name: name}
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
	offerIDs           []uuid.UUID
	cooperationTypeIDs []int64
	topicIDs           []int64
	contentFormatIDs   []int64
}

func collectOfferIDs(offers []*offerDomain.Offer) offerIDs {
	result := offerIDs{
		offerIDs:           make([]uuid.UUID, 0, len(offers)),
		cooperationTypeIDs: make([]int64, 0, len(offers)),
		topicIDs:           make([]int64, 0, len(offers)),
		contentFormatIDs:   make([]int64, 0, len(offers)),
	}

	for _, item := range offers {
		result.offerIDs = append(result.offerIDs, item.GetID())
		if item.GetCooperationTypeID() > 0 {
			result.cooperationTypeIDs = append(result.cooperationTypeIDs, item.GetCooperationTypeID())
		}
		if item.GetTopicID() > 0 {
			result.topicIDs = append(result.topicIDs, item.GetTopicID())
		}
		if item.GetContentFormatID() > 0 {
			result.contentFormatIDs = append(result.contentFormatIDs, item.GetContentFormatID())
		}
	}

	return result
}
