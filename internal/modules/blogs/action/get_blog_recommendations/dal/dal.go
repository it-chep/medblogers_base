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

func (r *Repository) GetBlogRecommendations(ctx context.Context, slug string) (blog.Blogs, error) {
	sql := `
		select b.id, b.name, b.slug, b.preview_text, b.created_at, b.ordering_number
		from blogs_recommendations br
		join blog src on br.blog_id = src.id
		join blog b on br.recommended_blog_id = b.id
		where src.slug = $1
		  and src.is_active is true
		  and b.is_active is true
		order by b.ordering_number
	`

	var blogs dao.BlogMiniatureDAOs
	err := pgxscan.Select(ctx, r.db, &blogs, sql, slug)
	if err != nil {
		return nil, err
	}

	return blogs.ToDomain(), nil
}

func (r *Repository) GetTopBlogsFallback(ctx context.Context, slug string) (blog.Blogs, error) {
	sql := `
		select id, name, slug, preview_text, created_at, ordering_number
		from blog
		where is_active is true
		  and slug <> $1
		order by ordering_number
		limit 3
	`

	var blogs dao.BlogMiniatureDAOs
	err := pgxscan.Select(ctx, r.db, &blogs, sql, slug)
	if err != nil {
		return nil, err
	}

	return blogs.ToDomain(), nil
}
