package dao

import (
	"database/sql"

	brandDomain "medblogers_base/internal/modules/admin/entities/promo_offers/domain/brand"
)

type BrandDAO struct {
	ID                 int64          `db:"id"`
	Photo              sql.NullString `db:"photo"`
	Title              sql.NullString `db:"title"`
	Slug               string         `db:"slug"`
	BusinessCategoryID sql.NullInt64  `db:"business_category_id"`
	Website            sql.NullString `db:"website"`
	Description        sql.NullString `db:"description"`
	IsActive           sql.NullBool   `db:"is_active"`
	CreatedAt          sql.NullTime   `db:"created_at"`
}

func (d BrandDAO) ToDomain() *brandDomain.Brand {
	var businessCategoryID int64
	if d.BusinessCategoryID.Valid {
		businessCategoryID = d.BusinessCategoryID.Int64
	}

	return brandDomain.New(
		brandDomain.WithID(d.ID),
		brandDomain.WithPhoto(d.Photo.String),
		brandDomain.WithTitle(d.Title.String),
		brandDomain.WithSlug(d.Slug),
		brandDomain.WithBusinessCategoryID(businessCategoryID),
		brandDomain.WithWebsite(d.Website.String),
		brandDomain.WithDescription(d.Description.String),
		brandDomain.WithIsActive(d.IsActive.Valid && d.IsActive.Bool),
		brandDomain.WithCreatedAt(&d.CreatedAt.Time),
	)
}

type Brands []BrandDAO

func (b Brands) ToDomain() brandDomain.Brands {
	items := make(brandDomain.Brands, 0, len(b))
	for _, item := range b {
		items = append(items, item.ToDomain())
	}

	return items
}
