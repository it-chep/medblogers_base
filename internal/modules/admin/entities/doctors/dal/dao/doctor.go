package dao

import (
	"database/sql"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"time"
)

// FullDoctorDAO карточка доктора для админки
type FullDoctorDAO struct {
	ID                int64          `json:"id" db:"id"`
	Name              string         `json:"name" db:"name"`
	Slug              string         `json:"slug" db:"slug"`
	Email             string         `json:"email" db:"email"`
	TgURL             sql.NullString `json:"tg_url" db:"tg_url"`
	InstURL           sql.NullString `json:"inst_url" db:"inst_url"`
	DzenURL           sql.NullString `json:"dzen_url" db:"dzen_url"`
	TgChannelURL      sql.NullString `json:"tg_channel_url" db:"tg_channel_url"`
	YoutubeURL        sql.NullString `json:"youtube_url" db:"youtube_url"`
	VkURL             sql.NullString `json:"vk_url" db:"vk_url"`
	SiteLink          sql.NullString `json:"prodoctorov" db:"prodoctorov"`
	TiktokURL         sql.NullString `json:"tiktok_url" db:"tiktok_url"`
	MainBlogTheme     sql.NullString `json:"main_blog_theme" db:"main_blog_theme"`
	CityID            int64          `json:"city_id" db:"city_id"`
	SpecialityID      int64          `json:"speciallity_id" db:"speciallity_id"`
	MedicalDirections sql.NullString `json:"medical_directions" db:"medical_directions"`
	DateCreated       time.Time      `json:"date_created" db:"date_created"`
	BirthDate         time.Time      `json:"birth_date" db:"birth_date"`
	S3Image           sql.NullString `json:"s3_image" db:"s3_image"`
	IsKFDoctor        sql.NullBool   `json:"is_kf_doctor" db:"is_kf_doctor"`
	CooperationType   sql.NullInt64  `json:"cooperation_type" db:"cooperation_type"`
	IsActive          bool           `json:"is_active" db:"is_active"`
}

// ToDomain конвертирует DAO в доменный объект
func (d FullDoctorDAO) ToDomain() *doctor.Doctor {
	return doctor.New(
		doctor.WithID(d.ID),
		doctor.WithName(d.Name),
		doctor.WithSlug(d.Slug),
		doctor.WithEmail(d.Email),
		doctor.WithTgURL(d.TgURL.String),
		doctor.WithInstURL(d.InstURL.String),
		doctor.WithDzenURL(d.DzenURL.String),
		doctor.WithTgChannelURL(d.TgChannelURL.String),
		doctor.WithYoutubeURL(d.YoutubeURL.String),
		doctor.WithVkURL(d.VkURL.String),
		doctor.WithSiteLink(d.SiteLink.String),
		doctor.WithTikTokURL(d.TiktokURL.String),
		doctor.WithMainBlogTheme(d.MainBlogTheme.String),
		doctor.WithMainSpecialityID(d.SpecialityID),
		doctor.WithMainCityID(d.CityID),
		doctor.WithCreatedDate(d.DateCreated),
		doctor.WithBirthDate(d.BirthDate),
		doctor.WithMedicalDirection(d.MedicalDirections.String),
		doctor.WithS3Image(doctor.S3Key(d.S3Image.String)),
		doctor.WithIsKFDoctor(d.IsKFDoctor.Bool),
		doctor.WithIsActive(d.IsActive),
		doctor.WithCooperationType(d.CooperationType.Int64),
	)
}

type DoctorMiniatureDao struct {
	ID              int64          `json:"id" db:"id"`
	Name            string         `json:"name" db:"name"`
	CooperationType sql.NullInt64  `json:"cooperation_type" db:"cooperation_type"`
	IsActive        bool           `json:"is_active" db:"is_active"`
	S3Image         sql.NullString `json:"s3_image" db:"s3_image"`
}

type DoctorMiniatureList []DoctorMiniatureDao

// ToDomain конвертирует DAO в доменный объект
func (d DoctorMiniatureDao) ToDomain() *doctor.Doctor {
	return doctor.New(
		doctor.WithID(d.ID),
		doctor.WithName(d.Name),
		doctor.WithIsActive(d.IsActive),
		doctor.WithCooperationType(d.CooperationType.Int64),
		doctor.WithS3Image(doctor.S3Key(d.S3Image.String)),
	)
}

func (d DoctorMiniatureList) ToDomain() []*doctor.Doctor {
	return lo.Map(d, func(item DoctorMiniatureDao, _ int) *doctor.Doctor {
		return item.ToDomain()
	})
}
