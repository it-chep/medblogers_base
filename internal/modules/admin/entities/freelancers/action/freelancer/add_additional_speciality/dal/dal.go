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

func (r *Repository) CreateCity(ctx context.Context, name string) error {
	sql := `
		insert into freelancers_city (name) values ($1)
		`

	_, err := r.db.Exec(ctx, sql, name)
	return err
}
