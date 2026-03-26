package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/samber/lo"

	"medblogers_base/internal/modules/admin/entities/promo_offers/dal/dao"
	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
	"medblogers_base/internal/modules/admin/entities/promo_offers/domain/dictionary"
	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetBrandByID(ctx context.Context, brandID int64) (*brandDomain.Brand, error) {
	sql := `
		select id, photo, title, slug, topic_id, website, description, is_active, created_at
		from brand
		where id = $1
	`

	var row dao.BrandDAO
	if err := pgxscan.Get(ctx, r.db, &row, sql, brandID); err != nil {
		return nil, err
	}

	return row.ToDomain(), nil
}

func (r *Repository) GetOfferByID(ctx context.Context, offerID uuid.UUID) (*offerDomain.Offer, error) {
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
		where id = $1
	`

	var row dao.OfferDAO
	if err := pgxscan.Get(ctx, r.db, &row, sql, offerID); err != nil {
		return nil, err
	}

	return row.ToDomain(), nil
}

func (r *Repository) GetTopicsByIDs(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		return map[int64]string{}, nil
	}

	sql := `
		select id, name
		from promo_offer_topic
		where id = any($1::bigint[])
	`

	var rows []dao.NamedDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql, pq.Int64Array(ids)); err != nil {
		return nil, err
	}

	result := make(map[int64]string, len(rows))
	for _, row := range rows {
		result[row.ID] = row.Name
	}

	return result, nil
}

func (r *Repository) GetBrandSocialNetworks(ctx context.Context, brandIDs []int64) (map[int64]dictionary.BrandSocialNetworks, error) {
	if len(brandIDs) == 0 {
		return map[int64]dictionary.BrandSocialNetworks{}, nil
	}

	sql := `
		select brand_id, social_network_id, link
		from brand_social_networks
		where brand_id = any($1::bigint[])
		order by brand_id, id
	`

	var rows []dao.BrandSocialNetworkLinkDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql, pq.Int64Array(brandIDs)); err != nil {
		return nil, err
	}

	networkIDs := lo.Uniq(lo.Map(rows, func(item dao.BrandSocialNetworkLinkDAO, _ int) int64 {
		return item.SocialNetworkID
	}))

	networksMap, err := r.getSocialNetworksByIDs(ctx, networkIDs)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]dictionary.BrandSocialNetworks, len(brandIDs))
	for _, row := range rows {
		network, ok := networksMap[row.SocialNetworkID]
		if !ok {
			continue
		}

		result[row.BrandID] = append(result[row.BrandID], dao.BrandSocialNetworkDAO{
			BrandID:         row.BrandID,
			SocialNetworkID: row.SocialNetworkID,
			Name:            network.Name,
			Slug:            network.Slug,
			Link:            row.Link,
		}.ToDomain())
	}

	return result, nil
}

func (r *Repository) getSocialNetworksByIDs(ctx context.Context, ids []int64) (map[int64]dao.SocialNetworkDAO, error) {
	if len(ids) == 0 {
		return map[int64]dao.SocialNetworkDAO{}, nil
	}

	sql := `
		select id, name, slug
		from social_networks
		where id = any($1::bigint[])
	`

	var rows []dao.SocialNetworkDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql, pq.Int64Array(ids)); err != nil {
		return nil, err
	}

	result := make(map[int64]dao.SocialNetworkDAO, len(rows))
	for _, row := range rows {
		result[row.ID] = row
	}

	return result, nil
}
