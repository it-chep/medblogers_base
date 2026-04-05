package dal

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) ActivateOffer(ctx context.Context, offerID uuid.UUID) error {
	_, err := r.db.Exec(ctx, `update promo_offer set is_active = true where id = $1`, offerID)
	return err
}
