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

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SearchBlogs(ctx context.Context, query string) (blog.Blogs, error) {
	sql := `
		select id, name, slug, preview_text, created_at, ordering_number
		from blog
		where is_active is true
		  and search_vector @@ websearch_to_tsquery('russian', $1)
		order by ts_rank(search_vector, websearch_to_tsquery('russian', $1)) desc,
		         ordering_number asc,
		         created_at desc
	`

	var blogs dao.BlogMiniatureDAOs
	err := pgxscan.Select(ctx, r.db, &blogs, sql, query)
	if err != nil {
		return nil, err
	}

	return blogs.ToDomain(), nil
}
