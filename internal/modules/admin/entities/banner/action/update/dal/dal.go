package dal

import (
	"context"

	"medblogers_base/internal/modules/admin/entities/banner/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpdateBanner(ctx context.Context, bannerID int64, req dto.UpdateRequest) error {
	sql := `
		update banners
		set
			name = $2,
			ordering_number = $3,
			banner_link = nullif($4, '')
		where id = $1
	`

	_, err := r.db.Exec(
		ctx,
		sql,
		bannerID,
		req.Name,
		req.OrderingNumber,
		req.BannerLink,
	)

	return err
}
