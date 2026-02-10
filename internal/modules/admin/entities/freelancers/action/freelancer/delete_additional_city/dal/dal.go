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

func (r *Repository) DeleteAdditionalCity(ctx context.Context, freelancerID, cityID int64) error {
	sql := `
		delete from freelancer_city_m2m where freelancer_id = $1 and city_id = $2
		`

	_, err := r.db.Exec(ctx, sql, freelancerID, cityID)
	return err
}
