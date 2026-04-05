package brand_seo

import (
	"context"

	actionDal "medblogers_base/internal/modules/promo_offers/action/brand_seo/dal"
	"medblogers_base/internal/modules/promo_offers/action/brand_seo/dto"
	"medblogers_base/internal/modules/promo_offers/client"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	"medblogers_base/internal/modules/promo_offers/service/image"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetBrandBySlug(ctx context.Context, slug string) (*brandDomain.Brand, error)
}

type Action struct {
	repository ActionDal
	image      *image.Service
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator) *Action {
	return &Action{
		repository: actionDal.NewRepository(pool),
		image:      image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, slug string) (*dto.Response, error) {
	logger.Message(ctx, "[PromoOffers][BrandSEO] Получение SEO бренда")

	brand, err := a.repository.GetBrandBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return &dto.Response{
		Title:       brand.GetTitle(),
		Description: brand.GetDescription(),
		ImageURL:    a.image.EnrichPhotoByKey(brand.GetPhoto()),
	}, nil
}
