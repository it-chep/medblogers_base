package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/update_draft_blog/dto"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе со статьями
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

// UpdateBlog обновление статьи в базе
func (r *Repository) UpdateBlog(ctx context.Context, blogID uuid.UUID, req dto.Request) error {
	sql := `
		update blog set name = $2, 
			slug = $3,
			body = $4,
			preview_text = $5,
			society_preview = $6,
			additional_seo_text = $7,
			ordering_number = $8 
		where id = $1
	`

	args := []interface{}{
		blogID.String(),
		req.Name,
		req.Slug,
		req.Body,
		req.PreviewText,
		req.SocietyPreviewText,
		req.AdditionalSEOText,
		req.OrderingNumber,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
