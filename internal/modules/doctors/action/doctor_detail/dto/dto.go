package dto

import "medblogers_base/internal/modules/doctors/domain/doctor"

type DoctorDTO struct {
	Name string `json:"name"`
	Slug string `json:"slug"`

	// соц.сети
	InstURL      string `json:"inst_url"`
	VkURL        string `json:"vk_url"`
	DzenURL      string `json:"dzen_url"`
	TgURL        string `json:"tg_url"`
	TgChannelURL string `json:"tg_channel_url"`
	YoutubeURL   string `json:"youtube_url"`
	TiktokURL    string `json:"tiktok_url"`
	SiteLink     string `json:"prodoctorov"`

	Cities       []string `json:"cities"`       // Города
	Specialities []string `json:"specialities"` // Специальности // todo не забыть про основную и тут надо вложенный структуры чтобы были ссылки

	// Подписчики
	TgSubsCount         string `json:"tg_subs_count"`
	TgSubsCountText     string `json:"tg_subs_count_text"`
	TgLastUpdatedDate   string `json:"tg_last_updated_date"`
	InstSubsCount       string `json:"inst_subs_count"`
	InstSubsCountText   string `json:"inst_subs_count_text"`
	InstLastUpdatedDate string `json:"inst_last_updated_date"`
}

func New(doc *doctor.Doctor) DoctorDTO {
	return DoctorDTO{}
}
