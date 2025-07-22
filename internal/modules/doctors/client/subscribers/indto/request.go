package indto

// Абстракция для бизнес-логики

type SocialMedia string

// todo int64
const (
	Telegram  SocialMedia = "tg"
	Instagram SocialMedia = "inst"
)

type GetDoctorsByFilterRequest struct {
	// соц.сеть
	SocialMedia []SocialMedia
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
