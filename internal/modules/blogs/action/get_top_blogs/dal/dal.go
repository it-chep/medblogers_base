package dal

import (
	"context"
	"medblogers_base/internal/modules/blogs/dal/blogs/dao"
	"medblogers_base/internal/modules/blogs/domain/blog"
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

// GetTopBlogs получение топ статей
func (r *Repository) GetTopBlogs(ctx context.Context) (blog.Blogs, error) {
	sql := `
		select id, name, slug, preview_text, created_at, ordering_number from blog 
			where is_active is true 
		order by ordering_number 
		limit 3
	`

	var blogs dao.BlogMiniatureDAOs
	err := pgxscan.Select(ctx, r.db, &blogs, sql)
	if err != nil {
		return nil, err
	}

	return blogs.ToDomain(), nil
}
