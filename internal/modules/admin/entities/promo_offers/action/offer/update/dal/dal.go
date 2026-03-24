package dal

import (
	"context"

	"github.com/google/uuid"

	"medblogers_base/internal/modules/admin/entities/promo_offers/action/offer/update/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpdateOffer(ctx context.Context, offerID uuid.UUID, req dto.UpdateRequest) error {
	sql := `
		update promo_offer
		set cooperation_type_id = $2,
			topic_id = $3,
			title = $4,
			description = $5,
			price = $6,
			content_format_id = $7,
			brand_id = $8,
			publication_date = $9,
			ad_marking_responsible = $10,
			responses_capacity = $11
		where id = $1
	`

	_, err := r.db.Exec(ctx, sql,
		offerID,
		req.CooperationTypeID,
		req.TopicID,
		req.Title,
		req.Description,
		req.Price,
		req.ContentFormatID,
		req.BrandID,
		req.PublicationDate,
		req.AdMarkingResponsible,
		req.ResponsesCapacity,
	)
	return err
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
