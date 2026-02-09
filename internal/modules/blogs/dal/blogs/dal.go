package blogs

import (
	"context"
	"medblogers_base/internal/modules/blogs/dal/blogs/dao"
	"medblogers_base/internal/modules/blogs/domain/blog_photo"
	"medblogers_base/internal/modules/blogs/domain/category"
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

// GetBlogsCategories получение категорий статей
func (r *Repository) GetBlogsCategories(ctx context.Context, blogIDs []uuid.UUID) (map[uuid.UUID]category.Categories, error) {
	sql := `
		select m2m.blog_id, c.id, c.name, c.font_color, c.bg_color
		from blog_category c
			join m2m_blog_category m2m on c.id = m2m.category_id
		where m2m.blog_id = any($1)
	`

	blogIDsStr := lo.Map(blogIDs, func(item uuid.UUID, index int) string {
		return item.String()
	})

	var categories dao.Categories
	err := pgxscan.Select(ctx, r.db, &categories, sql, blogIDsStr)
	if err != nil {
		return nil, err
	}

	resMap := make(map[uuid.UUID]category.Categories, len(blogIDs))
	for _, categ := range categories {
		resMap[categ.BlogID] = append(resMap[categ.BlogID], categ.ToDomain())
	}

	return resMap, nil
}

// GetBlogCategories получение категорий статей
func (r *Repository) GetBlogCategories(ctx context.Context, blogID uuid.UUID) (category.Categories, error) {
	sql := `
		select m2m.blog_id, c.id, c.name, c.font_color, c.bg_color
		from blog_category c
			join m2m_blog_category m2m on c.id = m2m.category_id
		where m2m.blog_id = $1
	`

	var categories dao.Categories
	err := pgxscan.Select(ctx, r.db, &categories, sql, blogID)

	return categories.ToDomain(), err
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
