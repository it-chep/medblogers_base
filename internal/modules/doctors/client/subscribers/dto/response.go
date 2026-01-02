package dto

// Абстракция над HTTP

type FilterInfo struct {
	// Название соцсети
	Name string `json:"name"`
	// Слаг для фильтрации в сервисе подписчиков "tg", "inst"
	Slug string `json:"slug"`
}

type FilterInfoResponse struct {
	Messengers []FilterInfo `json:"messengers"`
}

type GetAllSubscribersInfoResponse struct {
	// количество подписчиков
	SubscribersCount string `json:"subscribers_count"`
	// текст "подписчика", "подписчиков", "подписчик"
	SubscribersCountText string `json:"subscribers_count_text"`
	// дата последнего обновления в сервисе
	LastUpdated string `json:"last_updated"`
}

type GetSubscribersByDoctorIDs struct {
	//ID доктора
	DoctorID int64 `json:"doctor_id"`
	//количество подписчиков
	TgSubsCount      string `json:"telegram_subs_count"`
	InstSubsCount    string `json:"instagram_subs_count"`
	YoutubeSubsCount string `json:"youtube_subs_count"`
	VkSubsCount      string `json:"vk_subs_count"`
	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText      string `json:"telegram_subs_text"`
	InstSubsCountText    string `json:"instagram_subs_text"`
	YoutubeSubsCountText string `json:"youtube_subs_text"`
	VkSubsCountText      string `json:"vk_subs_text"`
}

type GetSubscribersByDoctorIDsResponse struct {
	Data map[int64]GetSubscribersByDoctorIDs `json:"data"`
}

type GetDoctorSubscribersResponse struct {
	//количество подписчиков
	TgSubsCount      string `json:"telegram_short"`
	InstSubsCount    string `json:"instagram_short"`
	YoutubeSubsCount string `json:"youtube_short"`
	VkSubsCount      string `json:"vk_short"`
	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText      string `json:"telegram_text"`
	InstSubsCountText    string `json:"instagram_text"`
	YoutubeSubsCountText string `json:"youtube_text"`
	VkSubsCountText      string `json:"vk_text"`
	//дата последнего обновления в сервисе
	TgLastUpdatedDate      string `json:"tg_last_updated_date"`
	InstLastUpdatedDate    string `json:"instagram_last_updated_date"`
	YoutubeLastUpdatedDate string `json:"youtube_last_updated_date"`
	VkLastUpdatedDate      string `json:"vk_last_updated_date"`
}

type GetDoctorsByFilter struct {
	//ID доктора
	DoctorID int64 `json:"doctor_id"`

	//количество подписчиков
	TgSubsCount      string `json:"telegram_short"`
	InstSubsCount    string `json:"inst_short"`
	YouTubeSubsCount string `json:"youtube_short"`
	VkSubsCount      string `json:"vk_short"`

	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText      string `json:"telegram_text"`
	InstSubsCountText    string `json:"inst_text"`
	YouTubeSubsCountText string `json:"youtube_text"`
	VkSubsCountText      string `json:"vk_text"`
}

type GetDoctorByFilter struct {
	Doctor GetDoctorsByFilter `json:"doctor"`
}

type GetDoctorsByFilterResponse struct {
	Doctors          []GetDoctorByFilter `json:"doctors"`
	DoctorsCount     int64               `json:"filtered_doctors_count"`
	SubscribersCount string              `json:"filtered_doctors_subscribers_count"`
}

type CheckTelegramInBlackListResponse struct {
	IsInBlackList bool `json:"is_in_blacklist"`
}
