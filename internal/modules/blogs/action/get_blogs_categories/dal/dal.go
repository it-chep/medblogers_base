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

func (r *Repository) AllBlogsCount(ctx context.Context) (int64, error) {
	sql := `
		select count(*) from blog where is_active is true
	`

	var count int64
	err := pgxscan.Get(ctx, r.db, &count, sql)
	if err != nil {
		return 0, err
	}
	return count, nil
}

type categoryCount struct {
	CategoryID int64 `db:"category_id"`
	Count      int64 `db:"count"`
}

func (r *Repository) CategoryBlogsCount(ctx context.Context) (map[int64]int64, error) {
	var (
		categories  []categoryCount
		categoryMap = make(map[int64]int64, len(categories))
		sql         = `
			select m2m.category_id, count(*) as count 
			from m2m_blog_category m2m
				join blog b on m2m.blog_id = b.id
			where b.is_active is true
			group by m2m.category_id
		`
	)

	err := pgxscan.Select(ctx, r.db, &categories, sql)
	if err != nil {
		return categoryMap, err
	}

	for _, c := range categories {
		categoryMap[c.CategoryID] = c.Count
	}

	return categoryMap, nil
}
