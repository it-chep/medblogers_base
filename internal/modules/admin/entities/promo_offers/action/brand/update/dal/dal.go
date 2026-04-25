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
			business_category_id = $3,
			website = $4,
			description = $5,
			about = $6
		where id = $1
	`

	_, err := r.db.Exec(ctx, sql,
		brandID,                // $1
		req.Title,              // $2
		req.BusinessCategoryID, // $3
		req.Website,            // $4
		req.Description,        // $5
		req.About,              // $6
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

func (r *Repository) UpdateBreadcrumb(ctx context.Context, slug, name string) error {
	sql := `
		update breadcrumbs
		set name = $2
		where url = '/brands/' || $1
	`

	_, err := r.db.Exec(ctx, sql, slug, name)
	return err
}
