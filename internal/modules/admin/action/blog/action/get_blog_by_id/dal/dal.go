package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/get_blog_by_id/dto"
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
