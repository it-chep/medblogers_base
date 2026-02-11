package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/blogs/dal/blogs/dao"
	"medblogers_base/internal/modules/blogs/domain/category"
	"medblogers_base/internal/pkg/postgres"
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

// GetAllCategories получение списка категорий
func (r *Repository) GetAllCategories(ctx context.Context) (category.Categories, error) {
	sql := `
		select 
		    id, name, font_color, bg_color 
		from blog_category
		order by id
	`

	var categories dao.Categories
	err := pgxscan.Select(ctx, r.db, &categories, sql)
	if err != nil {
		return nil, err
	}

	return categories.ToDomain(), nil
}
