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
	sql := `
		select id,
		       name,
		       created_at,
		       slug,
		       body,
		       is_active,
		       preview_text,
		       society_preview,
		       additional_seo_text,
		       ordering_number,
		       doctor_id
		from blog
		where id = $1
	`

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
			body = $3,
			preview_text = $4,
			society_preview = $5,
			additional_seo_text = $6,
			ordering_number = $7,
			doctor_id = $8,
			search_text = $9,
			search_vector =
				setweight(to_tsvector('russian', coalesce($2, '')), 'A') ||
				setweight(to_tsvector('russian', coalesce($5, '')), 'B') ||
				setweight(to_tsvector('russian', coalesce($10, '')), 'C')
		where id = $1
	`

	args := []interface{}{
		blogID.String(),        // $1
		req.Name,               // $2
		req.Body,               // $3
		req.PreviewText,        // $4
		req.SocietyPreviewText, // $5
		req.AdditionalSEOText,  // $6
		req.OrderingNumber,     // $7
	}

	if req.DoctorID != 0 {
		args = append(args, req.DoctorID) // $8
	} else {
		args = append(args, nil) // $8
	}

	args = append(args, req.SearchText) // $9

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateBreadcrumb(ctx context.Context, slug, name string) error {
	sql := `
		update breadcrumbs
		set name = $2
		where url = '/blogs/' || $1
	`

	_, err := r.db.Exec(ctx, sql, slug, name)
	return err
}
