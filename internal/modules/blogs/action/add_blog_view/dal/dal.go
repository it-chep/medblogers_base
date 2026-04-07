package dal

import (
	"context"
	"medblogers_base/internal/modules/blogs/dal/blogs/dao"
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе со статьями.
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// GetBlogBySlug получение статьи по slug.
func (r *Repository) GetBlogBySlug(ctx context.Context, slug string) (*blog.Blog, error) {
	sql := `
		select id, name, slug, body, preview_text, society_preview, additional_seo_text, created_at, ordering_number, doctor_id
		from blog
		where slug = $1
		  and is_active is true
	`

	var blogDAO dao.BlogDAO
	err := pgxscan.Get(ctx, r.db, &blogDAO, sql, slug)
	if err != nil {
		return nil, err
	}

	return blogDAO.ToDomain(), nil
}

// IsBlogViewedByCookieForLast7Days проверяет, был ли просмотр статьи этим cookie за последние 7 дней.
func (r *Repository) IsBlogViewedByCookieForLast7Days(ctx context.Context, blogID uuid.UUID, cookieID string) (bool, error) {
	sql := `
		select exists(
			select 1
			from blogs_views
			where blog_uuid = $1
			  and cookie_id = $2
			  and created_at between now() - interval '7 days' and now()
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, sql, blogID, cookieID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// CreateBlogView создает запись о просмотре статьи.
func (r *Repository) CreateBlogView(ctx context.Context, id uuid.UUID, blogID uuid.UUID, cookieID string) error {
	sql := `
		insert into blogs_views (id, blog_uuid, cookie_id)
		values ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, sql, id, blogID, cookieID)
	return err
}
