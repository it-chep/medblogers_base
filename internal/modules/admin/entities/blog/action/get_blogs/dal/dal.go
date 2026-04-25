package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blogs/dto"
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

// GetBlogs получение всех статей
func (r *Repository) GetBlogs(ctx context.Context) (dto.Blogs, error) {
	sql := `select id, name, is_active, ordering_number from blog order by ordering_number`

	var blogs dto.Blogs
	err := pgxscan.Select(ctx, r.db, &blogs, sql)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

// GetBlogsViewsCount получение количества просмотров для списка статей.
func (r *Repository) GetBlogsViewsCount(ctx context.Context, blogIDs []uuid.UUID) (map[uuid.UUID]int64, error) {
	if len(blogIDs) == 0 {
		return map[uuid.UUID]int64{}, nil
	}

	sql := `
		select blog_uuid, count(*) as views_count
		from blogs_views
		where blog_uuid = any($1::uuid[])
		group by blog_uuid
	`

	type blogViewsDAO struct {
		BlogUUID   uuid.UUID `db:"blog_uuid"`
		ViewsCount int64     `db:"views_count"`
	}

	ids := lo.Map(blogIDs, func(item uuid.UUID, _ int) string {
		return item.String()
	})

	var views []blogViewsDAO
	err := pgxscan.Select(ctx, r.db, &views, sql, ids)
	if err != nil {
		return nil, err
	}

	viewsMap := make(map[uuid.UUID]int64, len(blogIDs))
	for _, item := range views {
		viewsMap[item.BlogUUID] = item.ViewsCount
	}

	return viewsMap, nil
}
