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
		set photo = $2,
			title = $3,
			slug = $4,
			topic_id = $5,
			website = $6,
			description = $7
		where id = $1
	`

	_, err := r.db.Exec(ctx, sql,
		brandID,
		req.Photo,
		req.Title,
		req.Slug,
		req.TopicID,
		req.Website,
		req.Description,
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
