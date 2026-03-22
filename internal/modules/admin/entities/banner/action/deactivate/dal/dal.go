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

func (r *Repository) DeactivateBanner(ctx context.Context, bannerID int64) error {
	sql := `update banners set is_active = false where id = $1`
	_, err := r.db.Exec(ctx, sql, bannerID)
	return err
}
