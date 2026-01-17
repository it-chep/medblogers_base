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

// DeleteCategory удаление тега у статьи
func (r *Repository) DeleteCategory(ctx context.Context, blogID uuid.UUID, categoryID int64) error {
	sql := `
		delete from m2m_blog_category 
			where blog_id = $1 and category_id = $2
   	`

	_, err := r.db.Exec(ctx, sql, blogID, categoryID)
	return err
}
