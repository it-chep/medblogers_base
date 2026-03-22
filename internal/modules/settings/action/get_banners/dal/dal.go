package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"medblogers_base/internal/modules/settings/action/get_banners/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetActiveBanners(ctx context.Context) ([]dto.Banner, error) {
	sql := `
		select
			id,
			name,
			ordering_number,
			desktop_image::text as desktop_image,
			desktop_file_type,
			mobile_image::text as mobile_image,
			mobile_file_type,
			banner_link
		from banners
		where is_active is true
		order by ordering_number
	`

	var banners []dto.Banner
	err := pgxscan.Select(ctx, r.db, &banners, sql)

	return banners, err
}
