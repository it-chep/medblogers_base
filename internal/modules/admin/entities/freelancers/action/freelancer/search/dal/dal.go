package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/entities/freelancers/dal/dao"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
	"strings"
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
	rawQuery := strings.TrimSpace(query)
	pattern := "%" + rawQuery + "%"

	sql := `
		select id, name, is_active
		from freelancer f
		where f.name ilike $1
		   or exists(
		       select 1
		       from freelancers_price_list fpl
		       where fpl.freelancer_id = f.id
		         and fpl.search_vector @@ websearch_to_tsquery('russian', $2)
		   )
		order by f.name
		`

	var frlancers dao.MiniatureList
	err := pgxscan.Select(ctx, r.db, &frlancers, sql, pattern, rawQuery)

	return frlancers.ToDomain(), err
}
