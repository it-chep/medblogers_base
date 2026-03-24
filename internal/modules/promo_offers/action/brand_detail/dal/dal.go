package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetBrandBySlug(ctx context.Context, slug string) (*brandDomain.Brand, error) {
	sql := `
		select id, photo, title, slug, topic_id, website, description, is_active, created_at
		from brand
		where slug = $1 and is_active is true
	`

	var row dao.BrandDAO
	if err := pgxscan.Get(ctx, r.db, &row, sql, slug); err != nil {
		return nil, err
	}

	return row.ToDomain(), nil
}
