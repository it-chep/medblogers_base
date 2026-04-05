package dal

import (
	"context"
	"fmt"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/lib/pq"

	filterDTO "medblogers_base/internal/modules/promo_offers/action/filter_offers/dto"
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

func (r *Repository) FilterOffers(ctx context.Context, filter filterDTO.OfferFilter) ([]*offerDomain.Offer, error) {
	whereStmt, phValues := buildOfferFilterWhere(filter)

	sql := `
		select id, cooperation_type_id, business_category_id, description, brand_id, created_at
		from promo_offer o
		where o.is_active is true
	` + whereStmt + `
		order by o.created_at desc, o.id desc
	`

	var rows []dao.FilterOfferDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql, phValues...); err != nil {
		return nil, err
	}

	result := make([]*offerDomain.Offer, 0, len(rows))
	for _, row := range rows {
		result = append(result, row.ToDomain())
	}

	return result, nil
}

func buildOfferFilterWhere(filter filterDTO.OfferFilter) (string, []any) {
	var (
		builder  strings.Builder
		phValues []any
		ph       = 1
	)

	if len(filter.CooperationTypeIDs) > 0 {
		builder.WriteString(fmt.Sprintf(`
			and o.cooperation_type_id = any($%d::bigint[])
		`, ph))
		phValues = append(phValues, pq.Int64Array(filter.CooperationTypeIDs))
		ph++
	}

	return builder.String(), phValues
}
