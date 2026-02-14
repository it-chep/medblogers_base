package indto

// Абстракция для бизнес-логики

// GetDoctorsByFilterResponse ...
type GetDoctorsByFilterResponse struct {
	Doctors          map[int64]GetDoctorsByFilterDoctor
	DoctorsCount     int64
	SubscribersCount string
	OrderedIDs       []int64
}

// GetDoctorsByFilterDoctor ...
type GetDoctorsByFilterDoctor struct {
	//ID доктора
	DoctorID int64
	//количество подписчиков
	TgSubsCount      string
	InstSubsCount    string
	YouTubeSubsCount string
	VkSubsCount      string
	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText      string
	InstSubsCountText    string
	YouTubeSubsCountText string
	VkSubsCountText      string
}

type FilterInfoResponse struct {
	// Название соцсети
	Name string
	// Слаг для фильтрации в сервисе подписчиков "tg", "inst"
	Slug string
}

type GetAllSubscribersInfoResponse struct {
	// количество подписчиков
	SubscribersCount string
	// текст "подписчика", "подписчиков", "подписчик"
	SubscribersCountText string
	// дата последнего обновления в сервисе
	LastUpdated string
}

type GetSubscribersByDoctorIDsResponse struct {
	//ID доктора
	DoctorID int64
	//количество подписчиков
	TgSubsCount      string
	InstSubsCount    string
	YouTubeSubsCount string
	VkSubsCount      string
	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText      string
	InstSubsCountText    string
	YouTubeSubsCountText string
	VkSubsCountText      string
}

type GetDoctorSubscribersResponse struct {
	//количество подписчиков
	TgSubsCount          string
	InstSubsCount        string
	YouTubeSubsCountText string
	VkSubsCountText      string
	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText   string
	InstSubsCountText string
	YouTubeSubsCount  string
	VkSubsCount       string
	//дата последнего обновления в сервисе
	InstLastUpdatedDate    string
	TgLastUpdatedDate      string
	YouTubeLastUpdatedDate string
	VkLastUpdatedDate      string
}
