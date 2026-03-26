package save_brand_photo

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/save_brand_photo/dal"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/service/image"
	commonDal "medblogers_base/internal/modules/admin/entities/promo_offers/dal"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	SaveBrandPhoto(ctx context.Context, brandID int64, image string) error
}

type CommonDal interface {
	GetBrandByID(ctx context.Context, brandID int64) (*brandDomain.Brand, error)
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
	image     *image.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: commonDal.NewRepository(pool),
		actionDal: dal.NewRepository(pool),
		image:     image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, brandID int64, imageData []byte) (string, error) {
	item, err := a.commonDal.GetBrandByID(ctx, brandID)
	if err != nil {
		return "", err
	}

	filename, err := a.generateNewImageName(item, imageData)
	if err != nil {
		return "", err
	}

	newImageURL, err := a.image.SetImage(ctx, bytes.NewReader(imageData), filename)
	if err != nil {
		return "", err
	}

	if err = a.actionDal.SaveBrandPhoto(ctx, brandID, newImageURL); err != nil {
		return "", err
	}

	if item.GetPhoto() != "" {
		if err = a.image.DeleteImage(ctx, item.GetPhoto()); err != nil {
			return "", err
		}
	}

	return a.image.GetImageURL(newImageURL), nil
}

func (a *Action) generateNewImageName(item *brandDomain.Brand, imageData []byte) (string, error) {
	randUUID, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(imageData)
	fileType := strings.Split(contentType, "/")[1]

	return fmt.Sprintf("%s_%s.%s", item.GetSlug(), randUUID, fileType), nil
}
