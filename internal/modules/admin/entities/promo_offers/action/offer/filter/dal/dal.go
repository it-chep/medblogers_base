package dal

import (
	"context"
	"fmt"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/lib/pq"

	filterDTO "medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/filter/dto"
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

func (r *Repository) FilterOffers(ctx context.Context, req filterDTO.Request) (offerDomain.Offers, error) {
	whereStmt, phValues, nextPH := buildOffersWhere(req)

	sql := `
		select
			id,
			cooperation_type_id,
			business_category_id,
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
		where 1 = 1
	` + whereStmt + `
		order by created_at desc, id desc
	`
	// todo убрать where 1 = 1
	if req.Limit > 0 {
		sql += fmt.Sprintf(" limit $%d", nextPH)
		phValues = append(phValues, req.Limit)
		nextPH++
	}

	if req.Page > 0 && req.Limit > 0 {
		sql += fmt.Sprintf(" offset $%d", nextPH)
		phValues = append(phValues, (req.Page-1)*req.Limit)
	}

	var rows dao.Offers
	if err := pgxscan.Select(ctx, r.db, &rows, sql, phValues...); err != nil {
		return nil, err
	}

	return rows.ToDomain(), nil
}

func buildOffersWhere(req filterDTO.Request) (string, []any, int) {
	var (
		builder  strings.Builder
		phValues []any
		ph       = 1
	)

	if len(req.BrandIDs) > 0 {
		builder.WriteString(fmt.Sprintf(" and brand_id = any($%d::bigint[])", ph))
		phValues = append(phValues, pq.Int64Array(req.BrandIDs))
		ph++
	}

	if req.IsActive != nil {
		builder.WriteString(fmt.Sprintf(" and is_active = $%d", ph))
		phValues = append(phValues, *req.IsActive)
		ph++
	}

	return builder.String(), phValues, ph
}
