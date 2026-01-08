package dao

import (
	"database/sql"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/doctor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"time"
)

type FullFreelancerDAO struct {
	ID                   int64          `db:"id"`
	Email                string         `db:"email"`
	Slug                 string         `db:"slug"`
	Name                 string         `db:"name"`
	IsActive             bool           `db:"is_active"`
	TgUsername           sql.NullString `db:"tg_username"`
	PortfolioLink        sql.NullString `db:"portfolio_link"`
	SpecialityId         sql.NullInt64  `db:"speciality_id"`
	CityId               sql.NullInt64  `db:"city_id"`
	PriceCategory        sql.NullInt64  `db:"price_category"`
	S3Image              sql.NullString `db:"s3_image"`
	StartWorkingDate     *time.Time     `db:"start_working_date"`
	CooperationTypeId    sql.NullInt64  `db:"cooperation_type_id"`
	AgencyRepresentative sql.NullBool   `db:"agency_representative"`
}

func (d FullFreelancerDAO) ToDomain() *freelancer.Freelancer {
	return freelancer.New(
		freelancer.WithID(d.ID),
		freelancer.WithName(d.Name),
		freelancer.WithSlug(d.Slug),
		freelancer.WithIsActive(d.IsActive),
		freelancer.WithIsAgencyRepresentative(d.AgencyRepresentative.Bool),
		freelancer.WithMainSpecialityID(d.SpecialityId.Int64),
		freelancer.WithMainCityID(d.CityId.Int64),
		freelancer.WithPriceCategory(d.PriceCategory.Int64),
		freelancer.WithS3Image(d.S3Image.String),
		freelancer.WithCooperationTypeID(d.CooperationTypeId.Int64),
		freelancer.WithStartWorkingTime(*d.StartWorkingDate),
		freelancer.WithPortfolioLink(d.PortfolioLink.String),
		freelancer.WithTgURL(d.TgUsername.String),
	)
}

type FreelancerMiniatureDAO struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	IsActive bool   `db:"is_active"`
}

type MiniatureList []FreelancerMiniatureDAO

func (d FreelancerMiniatureDAO) ToDomain() *freelancer.Freelancer {
	return freelancer.New(
		freelancer.WithID(d.ID),
		freelancer.WithName(d.Name),
		freelancer.WithIsActive(d.IsActive),
	)
}

func (m MiniatureList) ToDomain() []*freelancer.Freelancer {
	return lo.Map(m, func(item FreelancerMiniatureDAO, _ int) *freelancer.Freelancer {
		return item.ToDomain()
	})
}

type RecommendationDoctorDAO struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (r *RecommendationDoctorDAO) ToDomain() *doctor.Doctor {
	return doctor.New(
		doctor.WithID(r.ID),
		doctor.WithName(r.Name),
	)
}

type Recommendations []RecommendationDoctorDAO

func (r Recommendations) ToDomain() []*doctor.Doctor {
	return lo.Map(r, func(item RecommendationDoctorDAO, _ int) *doctor.Doctor {
		return item.ToDomain()
	})
}

type PriceListDao struct {
	ID    int64          `db:"id"`
	Name  string         `db:"name"`
	Price pgtype.Numeric `db:"price"`
}
