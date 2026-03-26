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

func (r *Repository) DeleteNetwork(ctx context.Context, brandID, socialNetworkID int64) error {
	sql := `
		delete from brand_social_networks
		where brand_id = $1 and social_network_id = $2
	`

	_, err := r.db.Exec(ctx, sql, brandID, socialNetworkID)
	return err
}
