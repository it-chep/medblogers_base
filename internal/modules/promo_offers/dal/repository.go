package dal

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/lib/pq"

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

func (r *Repository) GetBrandsByIDs(ctx context.Context, ids []int64) (map[int64]*brandDomain.Brand, error) {
	if len(ids) == 0 {
		return map[int64]*brandDomain.Brand{}, nil
	}

	sql := `
		select id, photo, title, slug, business_category_id, website, description, about, is_active, created_at
		from brand
		where id = any($1::bigint[])
	`

	var rows []dao.BrandDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql, pq.Int64Array(ids)); err != nil {
		return nil, err
	}

	result := make(map[int64]*brandDomain.Brand, len(rows))
	for _, row := range rows {
		item := row.ToDomain()
		result[item.GetID()] = item
	}

	return result, nil
}

func (r *Repository) GetBusinessCategoriesByIDs(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		return map[int64]string{}, nil
	}

	sql := `
		select id, name
		from promo_offer_business_category
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

func (r *Repository) GetCooperationTypesByIDs(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		return map[int64]string{}, nil
	}

	sql := `
		select id, name
		from promo_offer_cooperation_type
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

func (r *Repository) GetContentFormatsByIDs(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		return map[int64]string{}, nil
	}

	sql := `
		select id, name
		from promo_offer_content_format
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

func (r *Repository) GetBrandSocialNetworks(ctx context.Context, brandIDs []int64) (map[int64][]dao.BrandSocialNetworkDAO, error) {
	if len(brandIDs) == 0 {
		return map[int64][]dao.BrandSocialNetworkDAO{}, nil
	}

	sql := `
		select bsn.brand_id, bsn.social_network_id, sn.name, sn.slug, bsn.link
		from brand_social_networks bsn
			join social_networks sn on sn.id = bsn.social_network_id
		where bsn.brand_id = any($1::bigint[])
		order by bsn.brand_id, bsn.id
	`

	var rows []dao.BrandSocialNetworkDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql, pq.Int64Array(brandIDs)); err != nil {
		return nil, err
	}

	result := make(map[int64][]dao.BrandSocialNetworkDAO, len(brandIDs))
	for _, row := range rows {
		result[row.BrandID] = append(result[row.BrandID], row)
	}

	return result, nil
}

func (r *Repository) GetOfferSocialNetworks(ctx context.Context, offerIDs []uuid.UUID) (map[uuid.UUID][]dao.OfferSocialNetworkDAO, error) {
	if len(offerIDs) == 0 {
		return map[uuid.UUID][]dao.OfferSocialNetworkDAO{}, nil
	}

	sql := `
		select posm.promo_offer_id, posm.social_network_id, sn.name, sn.slug
		from promo_offer_social_networks_m2m posm
			join social_networks sn on sn.id = posm.social_network_id
		where posm.promo_offer_id = any($1)
		order by posm.promo_offer_id, posm.id
	`

	uuids := make([]string, 0, len(offerIDs))
	for _, id := range offerIDs {
		uuids = append(uuids, id.String())
	}

	var rows []dao.OfferSocialNetworkDAO
	if err := pgxscan.Select(ctx, r.db, &rows, sql, uuids); err != nil {
		return nil, err
	}

	result := make(map[uuid.UUID][]dao.OfferSocialNetworkDAO, len(offerIDs))
	for _, row := range rows {
		result[row.OfferID] = append(result[row.OfferID], row)
	}

	return result, nil
}
