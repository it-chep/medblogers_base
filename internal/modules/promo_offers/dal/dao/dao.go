package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/promo_offers/domain/offer"
)

type BrandDAO struct {
	ID          int64          `db:"id"`
	Photo       sql.NullString `db:"photo"`
	Title       sql.NullString `db:"title"`
	Slug        string         `db:"slug"`
	TopicID     sql.NullInt64  `db:"topic_id"`
	Website     sql.NullString `db:"website"`
	Description sql.NullString `db:"description"`
	IsActive    sql.NullBool   `db:"is_active"`
	CreatedAt   sql.NullTime   `db:"created_at"`
}

func (d BrandDAO) ToDomain() *brandDomain.Brand {
	return brandDomain.New(
		brandDomain.WithID(d.ID),
		brandDomain.WithPhoto(d.Photo.String),
		brandDomain.WithTitle(d.Title.String),
		brandDomain.WithSlug(d.Slug),
		brandDomain.WithTopicID(nullInt64(d.TopicID)),
		brandDomain.WithWebsite(d.Website.String),
		brandDomain.WithDescription(d.Description.String),
		brandDomain.WithIsActive(d.IsActive.Valid && d.IsActive.Bool),
		brandDomain.WithCreatedAt(nullTime(d.CreatedAt)),
	)
}

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

type NamedDAO struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type SocialNetworkDAO struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Slug string `db:"slug"`
}

type BrandSocialNetworkLinkDAO struct {
	BrandID         int64  `db:"brand_id"`
	SocialNetworkID int64  `db:"social_network_id"`
	Link            string `db:"link"`
}

type BrandSocialNetworkDAO struct {
	BrandID         int64
	SocialNetworkID int64
	Name            string
	Slug            string
	Link            string
}

type OfferSocialNetworkLinkDAO struct {
	OfferID         uuid.UUID `db:"promo_offer_id"`
	SocialNetworkID int64     `db:"social_network_id"`
}

type OfferSocialNetworkDAO struct {
	OfferID         uuid.UUID
	SocialNetworkID int64
	Name            string
	Slug            string
}

type FilterCountDAO struct {
	ID          sql.NullInt64 `db:"id"`
	OffersCount int64         `db:"offers_count"`
}

func nullInt64(v sql.NullInt64) int64 {
	if !v.Valid {
		return 0
	}

	return v.Int64
}

func nullTime(v sql.NullTime) *time.Time {
	if !v.Valid {
		return nil
	}

	tm := v.Time
	return &tm
}
