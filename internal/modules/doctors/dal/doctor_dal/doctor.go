package doctor_dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
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

func (r Repository) GetDoctorsCount(ctx context.Context) (int64, error) {
	sql := `
		select count(*) as doctors_count
		from docstar_site_doctor d
		where d.is_active = true
	`

	var count int64
	if err := pgxscan.Get(ctx, r.db, &count, sql); err != nil {
		return 0, err
	}

	return count, nil
}
