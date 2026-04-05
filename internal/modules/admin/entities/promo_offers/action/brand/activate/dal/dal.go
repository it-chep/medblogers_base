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

func (r *Repository) ActivateBrand(ctx context.Context, brandID int64) error {
	_, err := r.db.Exec(ctx, `update brand set is_active = true where id = $1`, brandID)
	return err
}
