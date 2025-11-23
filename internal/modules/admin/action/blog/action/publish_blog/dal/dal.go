package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"medblogers_base/internal/modules/admin/action/blog/action/publish_blog/dto"
	"medblogers_base/internal/pkg/postgres"
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
