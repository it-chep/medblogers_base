package offer_seo

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	actionDal "medblogers_base/internal/modules/promo_offers/action/offer_seo/dal"
	"medblogers_base/internal/modules/promo_offers/action/offer_seo/dto"
	"medblogers_base/internal/modules/promo_offers/client"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/promo_offers/domain/offer"
	"medblogers_base/internal/modules/promo_offers/service/image"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetOfferByID(ctx context.Context, id uuid.UUID) (*offerDomain.Offer, error)
}

type CommonDal interface {
	GetBrandsByIDs(ctx context.Context, ids []int64) (map[int64]*brandDomain.Brand, error)
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

func (a *Action) Do(ctx context.Context, offerID uuid.UUID) (*dto.Response, error) {
	logger.Message(ctx, "[PromoOffers][OfferSEO] Получение SEO оффера")

	offer, err := a.repository.GetOfferByID(ctx, offerID)
	if err != nil {
		return nil, err
	}

	brandsMap, err := a.commonDal.GetBrandsByIDs(ctx, []int64{offer.GetBrandID()})
	if err != nil {
		return nil, err
	}

	brand, ok := brandsMap[offer.GetBrandID()]
	if !ok {
		return nil, errors.New("brand not found for promo offer seo")
	}

	brandTitle := brand.GetTitle()

	return &dto.Response{
		Title:       fmt.Sprintf("Рекламное предложение от %s", brandTitle),
		Description: fmt.Sprintf("Бренд %s ищет блогеров, чтобы реализовать свой проект: %s", brandTitle, offer.GetDescription()),
		ImageURL:    a.image.EnrichPhotoByKey(brand.GetPhoto()),
	}, nil
}
