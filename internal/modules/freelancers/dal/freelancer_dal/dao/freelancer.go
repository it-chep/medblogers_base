package dao

import "medblogers_base/internal/modules/freelancers/domain/freelancer"

type FreelancerDao struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	Slug         string `db:"slug"`
	S3Image      string `db:"s3_image"`
	IsActive     bool   `db:"is_active"`
	CityID       int64  `db:"city_id"`
	SpecialityID int64  `db:"speciallity_id"`
}

type FreelancerSeoInfo struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (f *FreelancerSeoInfo) ToDomain() *freelancer.Freelancer {
	return freelancer.New(
		freelancer.WithID(f.ID),
		freelancer.WithName(f.Name),
	)
}

type FreelancerMiniature struct {
	ID                       int64  `db:"id"`
	Name                     string `db:"name"`
	Slug                     string `db:"slug"`
	S3Image                  string `db:"s3_image"`
	PriceCategory            int64  `db:"price_category"`
	HasExperienceWithDoctors bool   `db:"is_worked_with_doctors"`
	CityName                 string `db:"city_name"`
	SpecialityName           string `db:"speciality_name"`
}

type FreelancerSearch struct {
	ID                       int64  `db:"id"`
	Name                     string `db:"name"`
	Slug                     string `db:"slug"`
	S3Image                  string `db:"s3_image"`
	PriceCategory            int64  `db:"price_category"`
	HasExperienceWithDoctors bool   `db:"is_worked_with_doctors"`
	CityName                 string `db:"city_name"`
	SpecialityName           string `db:"speciality_name"`
}

func (f *FreelancerSearch) ToDomain() *freelancer.Freelancer {
	return freelancer.New(
		freelancer.WithID(f.ID),
		freelancer.WithName(f.Name),
		freelancer.WithSlug(f.Slug),
		freelancer.WithPriceCategory(f.PriceCategory),
		freelancer.WithExperienceWithDoctors(f.HasExperienceWithDoctors),
		freelancer.WithS3Image(f.S3Image),
		// cityName
		// specaialityName
	)
}

type PriceCategory struct {
	ID               int64 `db:"id" json:"id"`
	FreelancersCount int64 `db:"freelancers_count" json:"freelancers_count"`
}
