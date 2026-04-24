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

func (r *Repository) AddPriceList(ctx context.Context, freelancerID int64, name string, price int64, priceTo *int64) error {
	sql := `
		insert into freelancers_price_list (freelancer_id, name, price, price_to, search_vector)
		values ($1, $2, $3, $4, to_tsvector('russian', coalesce($2, '')))
	`

	args := []interface{}{
		freelancerID,
		name,
		price,
		priceTo,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	return err
}
