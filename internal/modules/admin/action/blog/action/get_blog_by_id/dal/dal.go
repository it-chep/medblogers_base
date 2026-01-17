package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/get_blog_by_id/dto"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
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

// GetBlogByID получение статьи по ID
func (r *Repository) GetBlogByID(ctx context.Context, id uuid.UUID) (dto.Blog, error) {
	sql := `select * from blog where id = $1`

	var blog dto.Blog
	err := pgxscan.Get(ctx, r.db, &blog, sql, id.String())
	if err != nil {
		return dto.Blog{}, err
	}

	return blog, nil
}

// GetDoctorInfo информации о докторе
func (r *Repository) GetDoctorInfo(ctx context.Context, doctorID int64) (string, error) {
	sql := `select name from docstar_site_doctor where id = $1`

	var name string
	err := pgxscan.Get(ctx, r.db, &name, sql, doctorID)
	if err != nil {
		return "", err
	}

	return name, nil
}

// GetBlogCategories получение категорий статей
func (r *Repository) GetBlogCategories(ctx context.Context, blogID uuid.UUID) (dto.Categories, error) {
	sql := `
		select c.id, c.name 
		from blog_category c
		join m2m_blog_category m2m on c.id = m2m.category_id
		where m2m.blog_id = $1
	`

	var categories dto.Categories
	err := pgxscan.Select(ctx, r.db, &categories, sql, blogID.String())
	return categories, err
}
