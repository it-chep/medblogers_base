package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blogs/dto"
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

// GetBlogs получение всех статей
func (r *Repository) GetBlogs(ctx context.Context) (dto.Blogs, error) {
	sql := `select id, name, is_active, ordering_number from blog`

	var blogs dto.Blogs
	err := pgxscan.Select(ctx, r.db, &blogs, sql)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}
