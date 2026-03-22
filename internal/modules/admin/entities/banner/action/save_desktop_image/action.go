package save_desktop_image

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"medblogers_base/internal/modules/admin/client"
	saveDal "medblogers_base/internal/modules/admin/entities/banner/action/save_desktop_image/dal"
	"medblogers_base/internal/modules/admin/entities/banner/action/save_desktop_image/service/image"
	commonDal "medblogers_base/internal/modules/admin/entities/banner/dal"
	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type actionDal interface {
	SaveBannerDesktopImage(ctx context.Context, bannerID int64, imageID, fileType string) error
}

type commonGetter interface {
	GetBannerByID(ctx context.Context, bannerID int64) (*dto.Banner, error)
}

type Action struct {
	commonGetter commonGetter
	actionDal    actionDal
	image        *image.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		commonGetter: commonDal.NewRepository(pool),
		actionDal:    saveDal.NewRepository(pool),
		image:        image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, bannerID int64, imageData []byte) (string, error) {
	banner, err := a.commonGetter.GetBannerByID(ctx, bannerID)
	if err != nil {
		return "", err
	}

	imageID, filename, fileType, err := generateImageMeta(imageData)
	if err != nil {
		return "", err
	}

	imageKey, err := a.image.SetImage(ctx, bytes.NewReader(imageData), filename)
	if err != nil {
		return "", err
	}

	if err = a.actionDal.SaveBannerDesktopImage(ctx, bannerID, imageID, fileType); err != nil {
		return "", err
	}

	if banner.DesktopImage != "" && banner.DesktopFileType != "" {
		err = a.image.DeleteImage(ctx, fmt.Sprintf("%s.%s", banner.DesktopImage, banner.DesktopFileType))
		if err != nil {
			return "", err
		}
	}

	return a.image.GetImageURL(imageKey), nil
}

func generateImageMeta(imageData []byte) (imageID, filename, fileType string, err error) {
	uuidValue, err := uuid.NewV7()
	if err != nil {
		return "", "", "", err
	}

	contentType := http.DetectContentType(imageData)
	fileType = strings.Split(contentType, "/")[1]
	imageID = uuidValue.String()
	filename = fmt.Sprintf("%s.%s", imageID, fileType)

	return imageID, filename, fileType, nil
}
