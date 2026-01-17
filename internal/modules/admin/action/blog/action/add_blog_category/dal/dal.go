package dal

import (
	"context"
	"github.com/google/uuid"
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

// AddCategory добавление тега на статью
func (r *Repository) AddCategory(ctx context.Context, blogID uuid.UUID, categoryID int64) error {
	sql := `insert into m2m_blog_category (blog_id, category_id) values ($1, $2)`

	_, err := r.db.Exec(ctx, sql, blogID, categoryID)
	return err
}
