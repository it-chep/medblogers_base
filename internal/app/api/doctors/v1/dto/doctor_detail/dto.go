package doctor_detail

type CityItem struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SpecialityItem struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

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

	Cities       []CityItem       `json:"cities"`
	Specialities []SpecialityItem `json:"specialities"`

	MainCity       CityItem       `json:"main_city"`
	MainSpeciality SpecialityItem `json:"main_speciality"`

	// Подписчики
	TgSubsCount         string `json:"tg_subs_count"`
	TgSubsCountText     string `json:"tg_subs_count_text"`
	TgLastUpdatedDate   string `json:"tg_last_updated_date"`
	InstSubsCount       string `json:"inst_subs_count"`
	InstSubsCountText   string `json:"inst_subs_count_text"`
	InstLastUpdatedDate string `json:"inst_last_updated_date"`

	MainBlogTheme    string `json:"main_blog_theme"`
	MedicalDirection string `json:"medical_direction"`

	Image string `json:"image"`
}
