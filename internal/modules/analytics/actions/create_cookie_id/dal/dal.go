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

// CreateCookieUser создает анонимную сессию пользователя.
func (r *Repository) CreateCookieUser(ctx context.Context, cookieID uuid.UUID, domain string) error {
	sql := `
		insert into cookie_users (cookie_id, last_activity_at, domain_name)
		values ($1, now(), $2)
	`

	_, err := r.db.Exec(ctx, sql, cookieID, domain)
	return err
}
