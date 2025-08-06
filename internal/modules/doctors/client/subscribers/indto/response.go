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
	TgSubsCount string
	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText string
	//количество подписчиков
	InstSubsCount string
	//текст "подписчика", "подписчиков", "подписчик"
	InstSubsCountText string
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
	TgSubsCount string
	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText string
	//количество подписчиков
	InstSubsCount string
	//текст "подписчика", "подписчиков", "подписчик"
	InstSubsCountText string
}

type GetDoctorSubscribersResponse struct {
	//количество подписчиков
	TgSubsCount string
	//текст "подписчика", "подписчиков", "подписчик"
	TgSubsCountText string
	//дата последнего обновления в сервисе
	TgLastUpdatedDate string
	//количество подписчиков
	InstSubsCount string
	//текст "подписчика", "подписчиков", "подписчик"
	InstSubsCountText string
	//дата последнего обновления в сервисе
	InstLastUpdatedDate string
}
