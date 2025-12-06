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

// GetBlogDetail получение всей информации о статей
func (r *Repository) GetBlogDetail(ctx context.Context, slug string) (*blog.Blog, error) {
	sql := `
		select id, name, slug, body, preview_text, society_preview, additional_seo_text, created_at, ordering_number
		    from blog 
        where slug = $1 
          and is_active is true`

	var blogDAO dao.BlogDAO
	err := pgxscan.Get(ctx, r.db, &blogDAO, sql, slug)
	if err != nil {
		return nil, err
	}

	return blogDAO.ToDomain(), nil
}
