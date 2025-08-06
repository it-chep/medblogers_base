package dao

import (
	"database/sql"
	"medblogers_base/internal/modules/doctors/domain/doctor"
)

// DoctorDAO .
type DoctorDAO struct {
	ID                int64          `db:"id"`
	Name              string         `db:"name"`
	Slug              string         `db:"slug"`
	InstURL           sql.NullString `db:"inst_url"`
	VkURL             sql.NullString `db:"vk_url"`
	DzenURL           sql.NullString `db:"dzen_url"`
	TgURL             sql.NullString `db:"tg_url"`
	YoutubeURL        sql.NullString `db:"youtube_url"`
	Prodoctorov       sql.NullString `db:"prodoctorov"`
	TgChannelURL      sql.NullString `db:"tg_channel_url"`
	TiktokURL         sql.NullString `db:"tiktok_url"`
	S3Image           string         `db:"s3_image"`
	IsActive          bool           `db:"is_active"`
	MedicalDirections sql.NullString `db:"medical_directions"`
	MainBlogTheme     sql.NullString `db:"main_blog_theme"`
	CityID            int64          `db:"city_id"`
	SpecialityID      int64          `db:"speciallity_id"`
}

// ToDomain конвертирует DAO в доменный объект
func (d DoctorDAO) ToDomain() *doctor.Doctor {
	return doctor.New(
		doctor.WithID(d.ID),
		doctor.WithName(d.Name),
		doctor.WithSlug(d.Slug),
		doctor.WithS3Image(d.S3Image),
		doctor.WithTgURL(d.TgURL.String),
		doctor.WithInstURL(d.InstURL.String),
		doctor.WithDzenURL(d.DzenURL.String),
		doctor.WithTgChannelURL(d.TgChannelURL.String),
		doctor.WithYoutubeURL(d.YoutubeURL.String),
		doctor.WithVkURL(d.VkURL.String),
		doctor.WithTikTokURL(d.TiktokURL.String),
		doctor.WithMainBlogTheme(d.MainBlogTheme.String),
		doctor.WithMainSpecialityID(d.SpecialityID),
		doctor.WithMainCityID(d.CityID),
		doctor.WithMedicalDirection(d.MedicalDirections.String),
	)
}

// DoctorMiniatureDAO .
type DoctorMiniatureDAO struct {
	ID           int64          `db:"id"`
	Name         string         `db:"name"`
	Slug         string         `db:"slug"`
	InstURL      sql.NullString `db:"inst_url"`
	TgChannelURL sql.NullString `db:"tg_channel_url"`
	CityID       int64          `db:"city_id"`
	SpecialityID int64          `db:"speciallity_id"`
}

// ToDomain конвертирует DAO в доменный объект
func (d DoctorMiniatureDAO) ToDomain() *doctor.Doctor {
	return doctor.New(
		doctor.WithID(d.ID),
		doctor.WithSlug(d.Slug),
		doctor.WithName(d.Name),
		doctor.WithInstURL(d.InstURL.String),
		doctor.WithTgChannelURL(d.TgChannelURL.String),
		doctor.WithMainSpecialityID(d.SpecialityID),
		doctor.WithMainCityID(d.CityID),
	)
}

// DoctorSearchDAO .
type DoctorSearchDAO struct {
	ID             int64  `db:"id"`
	Name           string `db:"name"`
	Slug           string `db:"slug"`
	CityName       string `db:"city_name"`
	SpecialityName string `db:"speciality_name"`
}

// ToDomain конвертирует DAO в доменный объект
func (d DoctorSearchDAO) ToDomain() *doctor.Doctor {
	return doctor.New(
		doctor.WithID(d.ID),
		doctor.WithSlug(d.Slug),
		doctor.WithName(d.Name),
		doctor.WithCityName(d.CityName),
		doctor.WithSpecialityName(d.SpecialityName),
	)
}
