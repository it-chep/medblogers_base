package dal

import (
	"context"
	"fmt"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/lib/pq"

	filterDTO "medblogers_base/internal/modules/promo_offers/action/filter_brands/dto"
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

func (r *Repository) FilterBrands(ctx context.Context, filter filterDTO.BrandFilter) (brandDomain.Brands, error) {
	var (
		builder  strings.Builder
		phValues []any
		ph       = 1
		sql      = `
			select id, photo, title, slug, topic_id, website, description, is_active, created_at
			from brand b
			where b.is_active is true
		`
	)

	if len(filter.TopicIDs) > 0 {
		builder.WriteString(fmt.Sprintf(`
			and b.topic_id = any($%d::bigint[])
		`, ph))
		phValues = append(phValues, pq.Int64Array(filter.TopicIDs))
		ph++
	}

	if len(filter.SocialNetworkIDs) > 0 {
		builder.WriteString(fmt.Sprintf(`
			and exists (
				select 1
				from brand_social_networks bsn
				where bsn.brand_id = b.id
					and bsn.social_network_id = any($%d::bigint[])
			)
		`, ph))
		phValues = append(phValues, pq.Int64Array(filter.SocialNetworkIDs))
		ph++
	}

	builder.WriteString(`
		order by b.created_at desc, b.id desc
	`)

	var rows []dao.BrandDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql+builder.String(), phValues...); err != nil {
		return nil, err
	}

	result := make(brandDomain.Brands, 0, len(rows))
	for _, row := range rows {
		result = append(result, row.ToDomain())
	}

	return result, nil
}
