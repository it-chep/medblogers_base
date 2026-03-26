package dal

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/create/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateOffer(ctx context.Context, req dto.CreateRequest) (uuid.UUID, error) {
	sql := `
		insert into promo_offer (
			cooperation_type_id,
			business_category_id,
			title,
			description,
			price,
			content_format_id,
			brand_id,
			publication_date,
			ad_marking_responsible,
			responses_capacity
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		returning id
	`

	var offerID uuid.UUID
	err := r.db.QueryRow(ctx, sql,
		req.CooperationTypeID,
		req.BusinessCategoryID,
		req.Title,
		req.Description,
		req.Price,
		req.ContentFormatID,
		req.BrandID,
		req.PublicationDate,
		req.AdMarkingResponsible,
		req.ResponsesCapacity,
	).Scan(&offerID)

	return offerID, err
}

func (r *Repository) ReplaceOfferSocialNetworks(ctx context.Context, offerID uuid.UUID, socialNetworkIDs []int64) error {
	if _, err := r.db.Exec(ctx, `delete from promo_offer_social_networks_m2m where promo_offer_id = $1`, offerID); err != nil {
		return err
	}

	for _, socialNetworkID := range socialNetworkIDs {
		if _, err := r.db.Exec(
			ctx,
			`insert into promo_offer_social_networks_m2m (promo_offer_id, social_network_id) values ($1, $2)`,
			offerID,
			socialNetworkID,
		); err != nil {
			return err
		}
	}

	return nil
}
