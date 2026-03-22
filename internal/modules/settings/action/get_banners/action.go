package get_banners

import (
	"context"
	"database/sql"

	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/settings/action/get_banners/dal"
	"medblogers_base/internal/modules/settings/action/get_banners/dto"
	"medblogers_base/internal/modules/settings/action/get_banners/service/image"
	"medblogers_base/internal/pkg/postgres"
)

type bannersGetter interface {
	GetActiveBanners(ctx context.Context) ([]dto.Banner, error)
}

type Action struct {
	bannersGetter bannersGetter
	imageService  *image.Service
}

func New(pool postgres.PoolWrapper, cfg config.AppConfig) *Action {
	return &Action{
		bannersGetter: dal.NewRepository(pool),
		imageService:  image.New(cfg.GetSettingsBucket()),
	}
}

func (a *Action) Do(ctx context.Context) ([]dto.Banner, error) {
	banners, err := a.bannersGetter.GetActiveBanners(ctx)
	if err != nil {
		return nil, err
	}

	for i := range banners {
		banners[i].DesktopImage = toNullString(a.imageService.GetPhotoLink(
			banners[i].DesktopImage.String,
			banners[i].DesktopFileType.String,
		))
		banners[i].MobileImage = toNullString(a.imageService.GetPhotoLink(
			banners[i].MobileImage.String,
			banners[i].MobileFileType.String,
		))
	}

	return banners, nil
}

func toNullString(value string) sql.NullString {
	if value == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: value,
		Valid:  true,
	}
}
