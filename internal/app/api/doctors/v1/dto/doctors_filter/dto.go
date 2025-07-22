package doctors_filter

type DoctorItem struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`

	InstLink          string `json:"inst_link"`
	InstSubsCount     string `json:"inst_subs_count"`
	InstSubsCountText string `json:"inst_subs_count_text"`

	TgLink          string `json:"tg_link"`
	TgSubsCount     string `json:"tg_subs_count"`
	TgSubsCountText string `json:"tg_subs_count_text"`

	Speciality string `json:"speciality"` // Строка из основной и дополнительных специальностей
	City       string `json:"city"`       // Строка из основного и дополнительных городов

	Image string `json:"image"`
}

type Response struct {
	Doctors     []DoctorItem `json:"doctors"`
	Pages       int64        `json:"pages"`
	CurrentPage int64        `json:"page"`
}
