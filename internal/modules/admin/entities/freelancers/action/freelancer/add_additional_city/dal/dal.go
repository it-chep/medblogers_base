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

func (r *Repository) AddAdditionalCity(ctx context.Context, freelancerID, cityID int64) error {
	sql := `
		insert into freelancer_city_m2m (freelancer_id, city_id) values ($1,$2)
	`

	_, err := r.db.Exec(ctx, sql, freelancerID, cityID)
	return err
}
