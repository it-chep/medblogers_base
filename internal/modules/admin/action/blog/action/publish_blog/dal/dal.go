package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/publish_blog/dto"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
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

// GetBlogByID получение статьи по ID
func (r *Repository) GetBlogByID(ctx context.Context, id uuid.UUID) (dto.Blog, error) {
	sql := `select * from blog where id = $1`

	var blog dto.Blog
	err := pgxscan.Get(ctx, r.db, &blog, sql, id.String())
	if err != nil {
		return dto.Blog{}, err
	}

	return blog, nil
}

// PublishBlog публикация статьи
func (r *Repository) PublishBlog(ctx context.Context, id uuid.UUID) error {
	sql := `update blog set is_active = true where id = $1`

	_, err := r.db.Exec(ctx, sql, id.String())
	if err != nil {
		return err
	}

	return nil
}

// UnMarkBlogImagesIsPrimary убираем старую метку
func (r *Repository) UnMarkBlogImagesIsPrimary(ctx context.Context, blogID uuid.UUID) error {
	sql := `update blog_photos set is_primary = false where blog_id = $1 and is_primary is true`

	_, err := r.db.Exec(ctx, sql, blogID.String())
	if err != nil {
		return err
	}
	return nil
}

// MarkImageIsPrimary помечаем фотографию самой первой
func (r *Repository) MarkImageIsPrimary(ctx context.Context, imageID uuid.UUID) error {
	sql := `update blog_photos set is_primary = true where id = $1`

	_, err := r.db.Exec(ctx, sql, imageID.String())
	if err != nil {
		return err
	}
	return nil
}
