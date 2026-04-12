package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetBlogByID(ctx context.Context, blogID uuid.UUID) error {
	sql := `select id from blog where id = $1`

	var id uuid.UUID
	return pgxscan.Get(ctx, r.db, &id, sql, blogID.String())
}
