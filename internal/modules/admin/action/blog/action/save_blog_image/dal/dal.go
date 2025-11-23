package dal

import (
	"context"
	"github.com/google/uuid"
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

// SaveBlogPhoto сохраняет инфо о фотографии в базе
func (r *Repository) SaveBlogPhoto(ctx context.Context, blogID, imageID uuid.UUID, contentType string) error {
	sql := `insert into blog_photos (id, blog_id, file_type) values ($1, $2, $3)`

	args := []interface{}{
		imageID,
		blogID,
		contentType,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
