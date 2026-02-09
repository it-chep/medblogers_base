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

func (r *Repository) DeactivateFreelancer(ctx context.Context, freelancerID int64) (err error) {
	sql := `
		update freelancer set is_active = false where id = $1
	`

	_, err = r.db.Exec(ctx, sql, freelancerID)
	return err
}
