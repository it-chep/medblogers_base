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
func (r *Repository) CreateDraftBlogs(ctx context.Context, title, slug string, id uuid.UUID) error {
	sql := `insert into blog (id, name, slug, ordering_number) values ($1, $2, $3, 999)`

	args := []interface{}{
		id.String(),
		title,
		slug,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateBlogBreadcrumb(ctx context.Context, title, slug string) error {
	sql := `
		insert into breadcrumbs (name, url, parent_id)
		values (
			$1,
			'/blogs/' || $2,
			(select id from breadcrumbs where url = '/blogs' limit 1)
		)
		on conflict (url) do update
		set name = excluded.name,
			parent_id = excluded.parent_id
	`

	_, err := r.db.Exec(ctx, sql, title, slug)
	return err
}
