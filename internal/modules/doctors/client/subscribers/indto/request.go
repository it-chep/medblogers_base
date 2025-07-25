package indto

// Абстракция для бизнес-логики

type SocialMedia int64

const (
	Telegram SocialMedia = 1 + iota
	Instagram
)

func (sm SocialMedia) String() string {
	switch sm {
	case Telegram:
		return "tg"
	case Instagram:
		return "inst"
	}

	return ""
}

func NewSocialMedia(sm string) SocialMedia {
	switch sm {
	case "tg":
		return Telegram
	case "inst":
		return Instagram
	}

	return Telegram
}

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
