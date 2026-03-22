package dal

import (
	"context"

	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveBannerMobileImage(ctx context.Context, bannerID int64, imageID, fileType string) error {
	sql := `
		update banners
		set
			mobile_image = $2::uuid,
			mobile_file_type = $3
		where id = $1
	`

	_, err := r.db.Exec(ctx, sql, bannerID, imageID, fileType)
	return err
}
