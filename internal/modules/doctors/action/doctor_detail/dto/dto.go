package dto

import (
	"medblogers_base/internal/modules/doctors/domain/doctor"
)

// CityItem .
type CityItem struct {
	ID   int64
	Name string
}

// SpecialityItem .
type SpecialityItem struct {
	ID   int64
	Name string
}

// DoctorDTO представление доктора для детальной страницы
type DoctorDTO struct {
	Name string
	Slug string

	// соц.сети
	InstURL      string
	VkURL        string
	DzenURL      string
	TgURL        string
	TgChannelURL string
	YoutubeURL   string
	TiktokURL    string
	SiteLink     string

	Cities       []CityItem       // доп Города
	Specialities []SpecialityItem // доп Специальности

	// основной город
	MainCityID int64
	MainCity   CityItem
	// основная специальность
	MainSpecialityID int64
	MainSpeciality   SpecialityItem

	// Подписчики
	TgSubsCount      string
	InstSubsCount    string
	YouTubeSubsCount string

	TgSubsCountText      string
	InstSubsCountText    string
	YouTubeSubsCountText string

	TgLastUpdatedDate      string
	InstLastUpdatedDate    string
	YouTubeLastUpdatedDate string

	// фотка
	Image string

	MainBlogTheme    string
	MedicalDirection string

	IsKFDoctor bool
}

// New .
func New(doc *doctor.Doctor) *DoctorDTO {
	return &DoctorDTO{
		Name: doc.GetName(),
		Slug: doc.GetSlug(),

		InstURL:      doc.GetInstURL(),
		VkURL:        doc.GetVkURL(),
		DzenURL:      doc.GetDzenURL(),
		TgURL:        doc.GetTgURL(),
		TgChannelURL: doc.GetTgChannelURL(),
		TiktokURL:    doc.GetTiktokURL(),
		YoutubeURL:   doc.GetYoutubeURL(),
		SiteLink:     doc.GetSiteLink(),

		MainCityID:       int64(doc.GetMainCityID()),
		MainSpecialityID: int64(doc.GetMainSpecialityID()),

		MainBlogTheme:    doc.GetMainBlogTheme(),
		MedicalDirection: doc.GetMedicalDirection(),

		IsKFDoctor: doc.GetIsKFDoctor(),
	}
}
