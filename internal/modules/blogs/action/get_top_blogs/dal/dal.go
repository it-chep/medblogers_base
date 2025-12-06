package dal

import (
	"context"
	"medblogers_base/internal/modules/blogs/dal/blogs/dao"
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/modules/blogs/domain/blog_photo"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/samber/lo"
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

// GetTopBlogs получение топ статей
func (r *Repository) GetTopBlogs(ctx context.Context) (blog.Blogs, error) {
	sql := `
		select id, name, slug, preview_text, created_at, ordering_number from blog 
			where is_active is true 
		order by ordering_number 
		limit 3
	`

	var blogs dao.BlogMiniatureDAOs
	err := pgxscan.Select(ctx, r.db, &blogs, sql)
	if err != nil {
		return nil, err
	}

	return blogs.ToDomain(), nil
}

// GetPrimaryPhotos получение первых фотографий для миниатюр статей
func (r *Repository) GetPrimaryPhotos(ctx context.Context, blogIDs []uuid.UUID) (map[uuid.UUID]*blog_photo.BlogPhoto, error) {
	sql := `select id, blog_id, file_type, is_primary from blog_photos where is_primary is true and blog_id = any($1)`

	var photos dao.PrimaryPhotoDAOs
	ids := lo.Map(blogIDs, func(item uuid.UUID, _ int) string {
		return item.String()
	})
	err := pgxscan.Select(ctx, r.db, &photos, sql, ids)
	if err != nil {
		return nil, err
	}

	blogPhotoMap := lo.SliceToMap(photos, func(item *dao.PrimaryPhotoDAO) (uuid.UUID, *blog_photo.BlogPhoto) {
		return item.BlogID, item.ToDomain()
	})
	return blogPhotoMap, nil
}
