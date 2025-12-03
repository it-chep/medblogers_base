package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
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

// CreateDraftBlogs создание драфтовой статьи
func (r *Repository) CreateDraftBlogs(ctx context.Context, title string, id uuid.UUID) error {
	sql := `insert into blog (id, name, ordering_number) values ($1, $2, 999)`

	args := []interface{}{
		id.String(),
		title,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
