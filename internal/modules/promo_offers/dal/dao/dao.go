package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	brandDomain "medblogers_base/internal/modules/promo_offers/domain/brand"
	offerDomain "medblogers_base/internal/modules/promo_offers/domain/offer"
)

type BrandDAO struct {
	ID                 int64          `db:"id"`
	Photo              sql.NullString `db:"photo"`
	Title              sql.NullString `db:"title"`
	Slug               string         `db:"slug"`
	BusinessCategoryID sql.NullInt64  `db:"business_category_id"`
	Website            sql.NullString `db:"website"`
	Description        sql.NullString `db:"description"`
	About              sql.NullString `db:"about"`
	IsActive           sql.NullBool   `db:"is_active"`
	CreatedAt          sql.NullTime   `db:"created_at"`
}

func (d BrandDAO) ToDomain() *brandDomain.Brand {
	return brandDomain.New(
		brandDomain.WithID(d.ID),
		brandDomain.WithPhoto(d.Photo.String),
		brandDomain.WithTitle(d.Title.String),
		brandDomain.WithSlug(d.Slug),
		brandDomain.WithBusinessCategoryID(d.BusinessCategoryID.Int64),
		brandDomain.WithWebsite(d.Website.String),
		brandDomain.WithDescription(d.Description.String),
		brandDomain.WithAbout(d.About.String),
		brandDomain.WithIsActive(d.IsActive.Valid && d.IsActive.Bool),
		brandDomain.WithCreatedAt(&d.CreatedAt.Time),
	)
}

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
	var publicationDate *time.Time
	if d.PublicationDate.Valid {
		tm := d.PublicationDate.Time
		publicationDate = &tm
	}

	return offerDomain.New(
		offerDomain.WithID(d.ID),
		offerDomain.WithCooperationTypeID(d.CooperationTypeID.Int64),
		offerDomain.WithBusinessCategoryID(d.BusinessCategoryID.Int64),
		offerDomain.WithTitle(d.Title),
		offerDomain.WithDescription(d.Description.String),
		offerDomain.WithPrice(d.Price.Int64),
		offerDomain.WithContentFormatID(d.ContentFormatID.Int64),
		offerDomain.WithBrandID(d.BrandID.Int64),
		offerDomain.WithPublicationDate(publicationDate),
		offerDomain.WithAdMarkingResponsible(d.AdMarkingResponsible.String),
		offerDomain.WithResponsesCapacity(d.ResponsesCapacity.Int64),
		offerDomain.WithIsActive(d.IsActive.Valid && d.IsActive.Bool),
		offerDomain.WithCreatedAt(d.CreatedAt.Time),
	)
}

type FilterOfferDAO struct {
	ID                 uuid.UUID      `db:"id"`
	CooperationTypeID  sql.NullInt64  `db:"cooperation_type_id"`
	BusinessCategoryID sql.NullInt64  `db:"business_category_id"`
	Description        sql.NullString `db:"description"`
	BrandID            sql.NullInt64  `db:"brand_id"`
	CreatedAt          sql.NullTime   `db:"created_at"`
}

func (d FilterOfferDAO) ToDomain() *offerDomain.Offer {
	return offerDomain.New(
		offerDomain.WithID(d.ID),
		offerDomain.WithCooperationTypeID(d.CooperationTypeID.Int64),
		offerDomain.WithBusinessCategoryID(d.BusinessCategoryID.Int64),
		offerDomain.WithDescription(d.Description.String),
		offerDomain.WithBrandID(d.BrandID.Int64),
		offerDomain.WithCreatedAt(d.CreatedAt.Time),
	)
}

type OfferCardDAO struct {
	ID                uuid.UUID      `db:"id"`
	CooperationTypeID sql.NullInt64  `db:"cooperation_type_id"`
	Description       sql.NullString `db:"description"`
	Price             sql.NullInt64  `db:"price"`
	BrandID           sql.NullInt64  `db:"brand_id"`
	CreatedAt         sql.NullTime   `db:"created_at"`
}

func (d OfferCardDAO) ToDomain() *offerDomain.Offer {
	return offerDomain.New(
		offerDomain.WithID(d.ID),
		offerDomain.WithCooperationTypeID(d.CooperationTypeID.Int64),
		offerDomain.WithDescription(d.Description.String),
		offerDomain.WithPrice(d.Price.Int64),
		offerDomain.WithBrandID(d.BrandID.Int64),
		offerDomain.WithCreatedAt(d.CreatedAt.Time),
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
	OfferID         uuid.UUID `db:"promo_offer_id"`
	SocialNetworkID int64     `db:"social_network_id"`
	Name            string    `db:"name"`
	Slug            string    `db:"slug"`
}

type FilterCountDAO struct {
	ID          sql.NullInt64 `db:"id"`
	OffersCount int64         `db:"offers_count"`
}
