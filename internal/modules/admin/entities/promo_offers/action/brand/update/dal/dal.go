package dal

import (
	"context"

	"medblogers_base/internal/modules/admin/entities/promo_offers/action/brand/update/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpdateBrand(ctx context.Context, brandID int64, req dto.UpdateRequest) error {
	sql := `
		update brand
		set title = $2,
			slug = $3,
			business_category_id = $4,
			website = $5,
			description = $6,
			about = $7
		where id = $1
	`

	_, err := r.db.Exec(ctx, sql,
		brandID,
		req.Title,
		req.Slug,
		req.BusinessCategoryID,
		req.Website,
		req.Description,
		req.About,
	)
	return err
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
