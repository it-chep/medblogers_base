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

func (r *Repository) DeletePriceList(ctx context.Context, freelancerID, priceListID int64) error {
	sql := `
		delete from freelancers_price_list where id = $1 and freelancer_id = $2
	`

	_, err := r.db.Exec(ctx, sql, priceListID, freelancerID)
	return err
}
