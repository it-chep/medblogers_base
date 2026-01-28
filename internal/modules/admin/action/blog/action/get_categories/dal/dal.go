package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/get_categories/dto"
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

// GetCategories получение тегов для статей
func (r *Repository) GetCategories(ctx context.Context) (dto.Categories, error) {
	sql := `select id, name from blog_category`

	var categories dto.Categories
	err := pgxscan.Select(ctx, r.db, &categories, sql)

	return categories, err
}
