package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
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

func (r *Repository) SearchFreelancers(ctx context.Context, query string) ([]*freelancer.Freelancer, error) {
	sql := `
		select id, name, is_active from freelancer where name ilike $1
		`

	var frlancers dao.MiniatureList
	err := pgxscan.Select(ctx, r.db, &frlancers, sql, query)

	return frlancers.ToDomain(), err
}
