package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"medblogers_base/internal/modules/admin/entities/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetBrands(ctx context.Context) (brandDomain.Brands, error) {
	sql := `
		select id, photo, title, slug, business_category_id, website, description, is_active, created_at
		from brand
		order by id desc
	`

	var rows dao.Brands
	if err := pgxscan.Select(ctx, r.db, &rows, sql); err != nil {
		return nil, err
	}

	return rows.ToDomain(), nil
}
