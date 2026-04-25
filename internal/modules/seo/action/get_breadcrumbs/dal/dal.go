package dal

import (
	"context"
	"medblogers_base/internal/modules/seo/action/get_breadcrumbs/dto"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// todo возможно стоит сделать nested sets - https://habr.com/ru/articles/153861/ .
// todo Либо сделать кэш
// todo Либо сделать parent_ids []bigint если будет тормозить

// GetBreadcrumbs получение хлебных крошек
func (r *Repository) GetBreadcrumbs(ctx context.Context, path string) (dto.Breadcrumbs, error) {
	sql := `
		with recursive breadcrumbs_path as (
			select
				id,
				name,
				url,
				parent_id,
				array[id] as visited,
				1 as level
			from breadcrumbs
			where url = $1

			union all

			select
				b.id,
				b.name,
				b.url,
				b.parent_id,
				bp.visited || b.id,
				bp.level + 1
			from breadcrumbs b
				join breadcrumbs_path bp on bp.parent_id = b.id
			where not b.id = any(bp.visited)
		)
		select
			name,
			url as path
		from breadcrumbs_path
		order by level desc
	`

	var breadcrumbs dto.Breadcrumbs
	err := pgxscan.Select(ctx, r.db, &breadcrumbs, sql, path)
	if err != nil {
		return nil, err
	}

	return breadcrumbs, nil
}
