package dao

import (
	"database/sql"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
)

type FreelancerDao struct {
	ID           int64          `db:"id"`
	Name         string         `db:"name"`
	Slug         string         `db:"slug"`
	S3Image      sql.NullString `db:"s3_image"`
	IsActive     bool           `db:"is_active"`
	CityID       int64          `db:"city_id"`
	SpecialityID int64          `db:"speciallity_id"`
}

type FreelancerSeoInfo struct {
	ID           int64          `db:"id"`
	Name         string         `db:"name"`
	S3Image      sql.NullString `db:"s3_image"`
	StartWorking sql.NullTime   `db:"start_working_date"`
}

func (f *FreelancerSeoInfo) ToDomain() *freelancer.Freelancer {
	return freelancer.New(
		freelancer.WithID(f.ID),
		freelancer.WithName(f.Name),
		freelancer.WithS3Image(f.S3Image.String),
		freelancer.WithStartWorkingTime(f.StartWorking.Time),
	)
}

type FreelancerMiniature struct {
	ID                   int64          `db:"id"`
	Name                 string         `db:"name"`
	Slug                 string         `db:"slug"`
	S3Image              sql.NullString `db:"s3_image"`
	PriceCategory        int64          `db:"price_category"`
	AgencyRepresentative sql.NullBool   `db:"agency_representative"`
	SpecialityID         int64          `db:"speciality_id"`
	CityID               int64          `db:"city_id"`
}

func (m FreelancerMiniature) ToDomain() *freelancer.Freelancer {
	return freelancer.New(
		freelancer.WithID(m.ID),
		freelancer.WithName(m.Name),
		freelancer.WithSlug(m.Slug),
		freelancer.WithPriceCategory(m.PriceCategory),
		freelancer.WithS3Image(m.S3Image.String),
		freelancer.WithIsAgencyRepresentative(m.AgencyRepresentative.Bool),
		freelancer.WithMainCityID(m.CityID),
		freelancer.WithMainSpecialityID(m.SpecialityID),
	)
}

type Miniatures []FreelancerMiniature

func (m Miniatures) ToDomain() []*freelancer.Freelancer {
	domain := make([]*freelancer.Freelancer, 0, len(m))
	for _, miniature := range m {
		domain = append(domain, miniature.ToDomain())
	}
	return domain
}

type FreelancerSearch struct {
	ID                   int64          `db:"id"`
	Name                 string         `db:"name"`
	Slug                 string         `db:"slug"`
	S3Image              sql.NullString `db:"s3_image"`
	PriceCategory        int64          `db:"price_category"`
	AgencyRepresentative sql.NullBool   `db:"agency_representative"`
	CityName             string         `db:"city_name"`
	SpecialityName       string         `db:"speciality_name"`
}

func (f *FreelancerSearch) ToDomain() *freelancer.Freelancer {
	return freelancer.New(
		freelancer.WithID(f.ID),
		freelancer.WithName(f.Name),
		freelancer.WithSlug(f.Slug),
		freelancer.WithPriceCategory(f.PriceCategory),
		freelancer.WithS3Image(f.S3Image.String),
		freelancer.WithIsAgencyRepresentative(f.AgencyRepresentative.Bool),
		freelancer.WithCityName(f.CityName),
		freelancer.WithSpecialityName(f.SpecialityName),
	)
}

type PriceCategory struct {
	ID               int64 `db:"id" json:"id"`
	FreelancersCount int64 `db:"freelancers_count" json:"freelancers_count"`
}

// ----------------- //

type FreelancerDetail struct {
	ID                   int64          `db:"id"`
	Name                 string         `db:"name"`
	Slug                 string         `db:"slug"`
	TgUsername           string         `db:"tg_username"`
	PortfolioLink        sql.NullString `db:"portfolio_link"`
	SpecialityID         int64          `db:"speciality_id"`
	CityID               int64          `db:"city_id"`
	PriceCategory        int64          `db:"price_category"`
	S3Image              sql.NullString `db:"s3_image"`
	AgencyRepresentative sql.NullBool   `db:"agency_representative"`
	StartWorking         sql.NullTime   `db:"start_working_date"`
}

func (f FreelancerDetail) ToDomain() *freelancer.Freelancer {
	return freelancer.New(
		freelancer.WithID(f.ID),
		freelancer.WithName(f.Name),
		freelancer.WithSlug(f.Slug),
		freelancer.WithPriceCategory(f.PriceCategory),
		freelancer.WithTgURL(f.TgUsername),
		freelancer.WithPortfolioLink(f.PortfolioLink.String),
		freelancer.WithMainSpecialityID(f.SpecialityID),
		freelancer.WithMainCityID(f.CityID),
		freelancer.WithS3Image(f.S3Image.String),
		freelancer.WithIsAgencyRepresentative(f.AgencyRepresentative.Bool),
		freelancer.WithStartWorkingTime(f.StartWorking.Time),
	)
}
