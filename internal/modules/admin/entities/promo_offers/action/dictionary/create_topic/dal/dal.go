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

func (r *Repository) CreateTopic(ctx context.Context, name string) (int64, error) {
	var id int64
	err := r.db.QueryRow(ctx, `insert into promo_offer_business_category (name) values ($1) returning id`, name).Scan(&id)
	return id, err
}
