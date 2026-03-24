package dao

import (
	"database/sql"

	"github.com/google/uuid"

	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
)

type OfferDAO struct {
	ID                   uuid.UUID      `db:"id"`
	CooperationTypeID    sql.NullInt64  `db:"cooperation_type_id"`
	TopicID              sql.NullInt64  `db:"topic_id"`
	Title                string         `db:"title"`
	Description          sql.NullString `db:"description"`
	Price                sql.NullInt64  `db:"price"`
	ContentFormatID      sql.NullInt64  `db:"content_format_id"`
	BrandID              sql.NullInt64  `db:"brand_id"`
	PublicationDate      sql.NullTime   `db:"publication_date"`
	AdMarkingResponsible sql.NullString `db:"ad_marking_responsible"`
	ResponsesCapacity    sql.NullInt64  `db:"responses_capacity"`
	IsActive             sql.NullBool   `db:"is_active"`
	CreatedAt            sql.NullTime   `db:"created_at"`
}

func (d OfferDAO) ToDomain() *offerDomain.Offer {
	return offerDomain.New(
		offerDomain.WithID(d.ID),
		offerDomain.WithCooperationTypeID(nullInt64(d.CooperationTypeID)),
		offerDomain.WithTopicID(nullInt64(d.TopicID)),
		offerDomain.WithTitle(d.Title),
		offerDomain.WithDescription(d.Description.String),
		offerDomain.WithPrice(nullInt64(d.Price)),
		offerDomain.WithContentFormatID(nullInt64(d.ContentFormatID)),
		offerDomain.WithBrandID(nullInt64(d.BrandID)),
		offerDomain.WithPublicationDate(nullTime(d.PublicationDate)),
		offerDomain.WithAdMarkingResponsible(d.AdMarkingResponsible.String),
		offerDomain.WithResponsesCapacity(nullInt64(d.ResponsesCapacity)),
		offerDomain.WithIsActive(d.IsActive.Valid && d.IsActive.Bool),
		offerDomain.WithCreatedAt(nullTime(d.CreatedAt)),
	)
}

type Offers []OfferDAO

func (o Offers) ToDomain() offerDomain.Offers {
	items := make(offerDomain.Offers, 0, len(o))
	for _, item := range o {
		items = append(items, item.ToDomain())
	}

	return items
}
