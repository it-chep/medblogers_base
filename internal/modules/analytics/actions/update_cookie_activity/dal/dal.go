package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository .
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// UpdateLastActivity обновляет время последней активности cookie пользователя.
func (r *Repository) UpdateLastActivity(ctx context.Context, cookieID uuid.UUID) error {
	sql := `
		update cookie_users
		set last_activity_at = now()
		where cookie_id = $1
	`

	_, err := r.db.Exec(ctx, sql, cookieID)
	return err
}
