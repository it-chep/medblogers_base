package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"medblogers_base/internal/modules/admin/entities/banner/dal/dao"
	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetBanners(ctx context.Context) ([]dto.Banner, error) {
	sql := `
		select
			id,
			name,
			is_active,
			ordering_number,
			desktop_image::text as desktop_image,
			desktop_file_type,
			mobile_image::text as mobile_image,
			mobile_file_type,
			banner_link
		from banners
		order by ordering_number, id desc
	`

	var banners dao.Banners
	err := pgxscan.Select(ctx, r.db, &banners, sql)
	if err != nil {
		return nil, err
	}

	return banners.ToDTO(), nil
}

func (r *Repository) GetBannerByID(ctx context.Context, bannerID int64) (*dto.Banner, error) {
	sql := `
		select
			id,
			name,
			is_active,
			ordering_number,
			desktop_image::text as desktop_image,
			desktop_file_type,
			mobile_image::text as mobile_image,
			mobile_file_type,
			banner_link
		from banners
		where id = $1
	`

	var banner dao.Banner
	err := pgxscan.Get(ctx, r.db, &banner, sql, bannerID)
	if err != nil {
		return nil, err
	}

	bannerDTO := banner.ToDTO()

	return &bannerDTO, nil
}
