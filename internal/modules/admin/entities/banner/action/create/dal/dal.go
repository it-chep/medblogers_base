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

func (r *Repository) CreateBanner(ctx context.Context, req dto.UpdateRequest) (int64, error) {
	sql := `
		insert into banners (
			is_active,
			name,
			ordering_number,
			banner_link
		) values (
			false,
			$1,
			$2,
			nullif($3, '')
		)
		returning id
	`

	var bannerID int64
	err := r.db.QueryRow(
		ctx,
		sql,
		req.Name,
		req.OrderingNumber,
		req.BannerLink,
	).Scan(&bannerID)

	return bannerID, err
}
