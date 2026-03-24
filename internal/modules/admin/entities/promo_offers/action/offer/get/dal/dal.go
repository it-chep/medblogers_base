package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"medblogers_base/internal/modules/admin/entities/promo_offers/dal/dao"
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetOffers(ctx context.Context) (offerDomain.Offers, error) {
	sql := `
		select
			id,
			cooperation_type_id,
			topic_id,
			title,
			description,
			price,
			content_format_id,
			brand_id,
			publication_date,
			ad_marking_responsible,
			responses_capacity,
			is_active,
			created_at
		from promo_offer
		order by created_at desc, id desc
	`

	var rows dao.Offers
	if err := pgxscan.Select(ctx, r.db, &rows, sql); err != nil {
		return nil, err
	}

	return rows.ToDomain(), nil
}
