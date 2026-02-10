package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AddNetwork(ctx context.Context, freelancerID, networkID int64) error {
	sql := `
		insert into freelancer_social_networks_m2m (social_network_id, freelancer_id) values ($1, $2)
		`

	_, err := r.db.Exec(ctx, sql, networkID, freelancerID)
	return err
}
