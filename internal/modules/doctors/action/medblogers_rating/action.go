package medblogers_rating

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/medblogers_rating/dal"
	"medblogers_base/internal/modules/doctors/action/medblogers_rating/dto"
	"medblogers_base/internal/modules/doctors/action/medblogers_rating/service/image"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	dal          *dal.Repository
	imageService *image.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		dal:          dal.NewRepository(pool),
		imageService: image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context) ([]dto.RatingItem, error) {
	items, err := a.dal.GetRating(ctx)
	if err != nil {
		return nil, err
	}

	a.imageService.EnrichImages(ctx, items)

	return items, nil
}
