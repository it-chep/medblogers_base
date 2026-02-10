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

func (r *Repository) AddPriceList(ctx context.Context, freelancerID int64, name string, price int64) error {
	sql := `
		insert into freelancers_price_list (freelancer_id, name, price) values ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, sql, freelancerID, name, price)
	return err
}
