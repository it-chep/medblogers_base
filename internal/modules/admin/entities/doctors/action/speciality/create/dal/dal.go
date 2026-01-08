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

func (r *Repository) CreateSpeciality(ctx context.Context, name string, isOnlyAdditional bool) error {
	sql := `
		insert into docstar_site_speciallity (name, is_only_additional) values ($1, $2)
		`

	_, err := r.db.Exec(ctx, sql, name, isOnlyAdditional)
	return err
}
