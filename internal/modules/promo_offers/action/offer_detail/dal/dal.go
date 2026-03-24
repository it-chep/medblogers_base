package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"

	"medblogers_base/internal/modules/promo_offers/dal/dao"
	offerDomain "medblogers_base/internal/modules/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetOfferByID(ctx context.Context, id uuid.UUID) (*offerDomain.Offer, error) {
	sql := `
		select id, cooperation_type_id, topic_id, title, description, price, content_format_id, brand_id,
		       publication_date, ad_marking_responsible, responses_capacity, is_active, created_at
		from promo_offer
		where id = $1 and is_active is true
	`

	var row dao.OfferDAO
	if err := pgxscan.Get(ctx, r.db, &row, sql, id); err != nil {
		return nil, err
	}

	return row.ToDomain(), nil
}
