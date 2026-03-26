package dal

import (
	"context"

	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/create/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateBrand(ctx context.Context, req dto.CreateRequest) (int64, error) {
	sql := `
		insert into brand (photo, title, slug, business_category_id, website, description)
		values ($1, $2, $3, $4, $5, $6)
		returning id
	`

	var brandID int64
	err := r.db.QueryRow(ctx, sql,
		req.Photo,
		req.Title,
		req.Slug,
		req.BusinessCategoryID,
		req.Website,
		req.Description,
	).Scan(&brandID)

	return brandID, err
}

func (r *Repository) ReplaceBrandSocialNetworks(ctx context.Context, brandID int64, items []dto.SocialNetworkInput) error {
	if _, err := r.db.Exec(ctx, `delete from brand_social_networks where brand_id = $1`, brandID); err != nil {
		return err
	}

	for _, item := range items {
		if _, err := r.db.Exec(
			ctx,
			`insert into brand_social_networks (brand_id, social_network_id, link) values ($1, $2, $3)`,
			brandID,
			item.SocialNetworkID,
			item.Link,
		); err != nil {
			return err
		}
	}

	return nil
}
