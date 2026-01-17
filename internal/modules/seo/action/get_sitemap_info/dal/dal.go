package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
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

func (r *Repository) GetAllDoctorsSlugs(ctx context.Context) ([]string, error) {
	sql := `
		select slug 
		from docstar_site_doctor 
        where is_active is true 
        order by id
	`

	var slugs []string
	err := pgxscan.Select(ctx, r.db, &slugs, sql)

	return slugs, err
}
func (r *Repository) GetAllFreelancersSlugs(ctx context.Context) ([]string, error) {
	sql := `
		select slug 
		from freelancer 
        where is_active is true 
        order by id
	`

	var slugs []string
	err := pgxscan.Select(ctx, r.db, &slugs, sql)

	return slugs, err
}

func (r *Repository) GetAllBlogsSlugs(ctx context.Context) ([]string, error) {
	sql := `
		select slug 
		from blog 
        where is_active is true 
        order by created_at
	`

	var slugs []string
	err := pgxscan.Select(ctx, r.db, &slugs, sql)

	return slugs, err
}
