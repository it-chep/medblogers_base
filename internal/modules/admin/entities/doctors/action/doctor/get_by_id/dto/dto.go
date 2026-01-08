package dto

import (
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"time"
)

type DoctorDTO struct {
	ID   int64
	Name string
	Slug string

	InstURL      string
	VkURL        string
	DzenURL      string
	TgURL        string
	TgChannelURL string
	YoutubeURL   string
	TikTokURL    string
	SiteLink     string

	S3Key doctor.S3Key

	MainCity       City
	MainSpeciality Speciality

	AdditionalCities       []City
	AdditionalSpecialities []Speciality

	MainBlogTheme string
	Image         string

	IsKfDoctor bool
	IsActive   bool

	BirthDate string
	CreatedAt string

	MedicalDirections    string
	MarketingPreferences string

	SubscribersInfo []Subscribers
	CooperationType doctor.CooperationType
}

type City struct {
	ID   int64
	Name string
}

type Speciality struct {
	ID   int64
	Name string
}

type Subscribers struct {
	Key             string
	SubsCount       string
	SubsCountText   string
	LastUpdatedDate string
}

func New(doc *doctor.Doctor) *DoctorDTO {
	return &DoctorDTO{
		ID:           int64(doc.GetID()),
		Name:         doc.GetName(),
		Slug:         doc.GetSlug(),
		InstURL:      doc.GetInstURL(),
		VkURL:        doc.GetVkURL(),
		DzenURL:      doc.GetDzenURL(),
		TgURL:        doc.GetTgURL(),
		TgChannelURL: doc.GetTgChannelURL(),
		YoutubeURL:   doc.GetYoutubeURL(),
		TikTokURL:    doc.GetTiktokURL(),
		SiteLink:     doc.GetSiteLink(),
		MainCity: City{
			ID: int64(doc.GetMainCityID()),
		},
		MainSpeciality: Speciality{
			ID: int64(doc.GetMainSpecialityID()),
		},
		MainBlogTheme:        doc.GetMainBlogTheme(),
		S3Key:                doc.GetS3Key(),
		IsKfDoctor:           doc.GetIsKFDoctor(),
		IsActive:             doc.GetIsActive(),
		BirthDate:            doc.GetBirthDate().Format(time.DateTime),
		CreatedAt:            doc.GetCreatedAt().Format(time.DateTime),
		MedicalDirections:    doc.GetMedicalDirection(),
		MarketingPreferences: doc.GetMarketingPreferences(),
		CooperationType:      doc.GetCooperationType(),
	}
}
