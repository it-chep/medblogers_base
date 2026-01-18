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

// ChangeMMActivity изменяет активность ММ
func (r *Repository) ChangeMMActivity(ctx context.Context, mmID int64, activity bool) error {
	sql := `update mm set is_active = $1 where id = $2`

	_, err := r.db.Exec(ctx, sql, activity, mmID)
	return err
}
