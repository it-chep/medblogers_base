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

func (r *Repository) GetBrandByOffer(ctx context.Context, offerID string) (*brandDomain.Brand, error) {
	sql := `
		select distinct b.id, b.photo, b.title, b.slug, b.business_category_id, b.website, b.description, b.about, b.is_active, b.created_at
		from brand b
			join promo_offer po on b.id = po.brand_id
		where po.id = $1 and b.is_active is true
	`

	var row dao.BrandDAO
	if err := pgxscan.Get(ctx, r.db, &row, sql, offerID); err != nil {
		return nil, err
	}

	return row.ToDomain(), nil
}
