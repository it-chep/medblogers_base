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

func (r *Repository) AddNetwork(ctx context.Context, brandID, socialNetworkID int64, link string) error {
	sql := `
		insert into brand_social_networks (brand_id, social_network_id, link)
		values ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, sql, brandID, socialNetworkID, link)
	return err
}
