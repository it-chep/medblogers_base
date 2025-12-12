package dal

import (
	"context"
	"medblogers_base/internal/modules/blogs/dal/blogs/dao"
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/modules/blogs/domain/blog_photo"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

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

// GetPrimaryPhoto получение первой фотографии для сеошки
func (r *Repository) GetPrimaryPhoto(ctx context.Context, blogID uuid.UUID) (*blog_photo.BlogPhoto, error) {
	sql := `
		select id, blog_id, file_type, is_primary 
		from blog_photos 
		where is_primary is true 
		  and blog_id = $1
	`

	var photo dao.PrimaryPhotoDAO

	err := pgxscan.Get(ctx, r.db, &photo, sql, blogID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return photo.ToDomain(), nil
}
