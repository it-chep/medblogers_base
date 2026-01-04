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

func (r *Repository) DeleteSpeciality(ctx context.Context, id int64) error {
	sql := `
		delete from freelancers_speciality where id = $1
		`

	_, err := r.db.Exec(ctx, sql, id)
	return err
}
