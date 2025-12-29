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

func (r *Repository) DeactivateDoctor(ctx context.Context, doctorID int64) (err error) {
	sql := `
		update docstar_site_doctor set is_active = false where id = $1
	`
	_, err = r.db.Exec(ctx, sql, doctorID)
	return err
}
