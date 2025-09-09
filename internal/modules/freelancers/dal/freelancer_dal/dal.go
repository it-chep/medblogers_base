package freelancer_dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
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

func (r Repository) GetFreelancersCount(ctx context.Context) (int64, error) {
	sql := `
		select count(*) as count
		from freelancer f
		where f.is_active = true
	`

	var count int64
	if err := pgxscan.Get(ctx, r.db, &count, sql); err != nil {
		return 0, err
	}

	return count, nil
}
