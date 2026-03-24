package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"medblogers_base/internal/modules/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/promo_offers/domain/offer"
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

func (r *Repository) GetOffersByBrandID(ctx context.Context, brandID int64) ([]*offerDomain.Offer, error) {
	sql := `
		select id, cooperation_type_id, topic_id, title, description, price, content_format_id, brand_id,
		       publication_date, ad_marking_responsible, responses_capacity, is_active, created_at
		from promo_offer
		where brand_id = $1 and is_active is true
		order by created_at desc, id desc
	`

	var rows []dao.OfferDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql, brandID); err != nil {
		return nil, err
	}

	result := make([]*offerDomain.Offer, 0, len(rows))
	for _, row := range rows {
		result = append(result, row.ToDomain())
	}

	return result, nil
}
