package indto

// Абстракция для бизнес-логики

type GetDoctorsByFilterRequest struct {
	// соц.сеть
	SocialMedia []string
	// офсет
	Offset int64
	// лимит
	Limit int64
	// максимальное количество подписчиков
	MaxSubscribers int64
	// минимальное количество подписчиков
	MinSubscribers int64
}

type CreateDoctorRequest struct {
	// username telegram
	Telegram string
	// username telegram
	Instagram string
}

type UpdateDoctorRequest struct {
	// username telegram
	Telegram string
	// username telegram
	Instagram string
}
