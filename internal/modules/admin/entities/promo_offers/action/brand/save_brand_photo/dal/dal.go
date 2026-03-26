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

func (r *Repository) SaveBrandPhoto(ctx context.Context, brandID int64, image string) error {
	_, err := r.db.Exec(ctx, `update brand set photo = $2 where id = $1`, brandID, image)
	return err
}
