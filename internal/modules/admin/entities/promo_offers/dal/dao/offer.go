package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	offerDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/offer"
)

type OfferDAO struct {
	ID                   uuid.UUID      `db:"id"`
	CooperationTypeID    sql.NullInt64  `db:"cooperation_type_id"`
	BusinessCategoryID   sql.NullInt64  `db:"business_category_id"`
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
	var cooperationTypeID int64
	if d.CooperationTypeID.Valid {
		cooperationTypeID = d.CooperationTypeID.Int64
	}

	var businessCategoryID int64
	if d.BusinessCategoryID.Valid {
		businessCategoryID = d.BusinessCategoryID.Int64
	}

	var price int64
	if d.Price.Valid {
		price = d.Price.Int64
	}

	var contentFormatID int64
	if d.ContentFormatID.Valid {
		contentFormatID = d.ContentFormatID.Int64
	}

	var brandID int64
	if d.BrandID.Valid {
		brandID = d.BrandID.Int64
	}

	var publicationDate *time.Time
	if d.PublicationDate.Valid {
		tm := d.PublicationDate.Time
		publicationDate = &tm
	}

	var responsesCapacity int64
	if d.ResponsesCapacity.Valid {
		responsesCapacity = d.ResponsesCapacity.Int64
	}

	return offerDomain.New(
		offerDomain.WithID(d.ID),
		offerDomain.WithCooperationTypeID(cooperationTypeID),
		offerDomain.WithBusinessCategoryID(businessCategoryID),
		offerDomain.WithTitle(d.Title),
		offerDomain.WithDescription(d.Description.String),
		offerDomain.WithPrice(price),
		offerDomain.WithContentFormatID(contentFormatID),
		offerDomain.WithBrandID(brandID),
		offerDomain.WithPublicationDate(publicationDate),
		offerDomain.WithAdMarkingResponsible(d.AdMarkingResponsible.String),
		offerDomain.WithResponsesCapacity(responsesCapacity),
		offerDomain.WithIsActive(d.IsActive.Valid && d.IsActive.Bool),
		offerDomain.WithCreatedAt(&d.CreatedAt.Time),
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
