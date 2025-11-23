package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"medblogers_base/internal/modules/admin/action/blog/action/delete_blog_image/dto"
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

// GetBlogImageByID получение фотографии по ID
func (r *Repository) GetBlogImageByID(ctx context.Context, imageID uuid.UUID) (dto.ImageDTO, error) {
	sql := `select * from blog_photos where id = $1`

	var image dto.ImageDTO
	err := pgxscan.Get(ctx, r.db, &image, sql, imageID)
	if err != nil {
		return image, err
	}

	return image, nil
}

// DeleteBlogImage удаление фотографии статьи
func (r *Repository) DeleteBlogImage(ctx context.Context, blogID, imageID uuid.UUID) error {
	sql := `delete from blog_photos where id = $1 and blog_id = $2`

	args := []interface{}{
		imageID,
		blogID,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
