package indto

// Абстракция для бизнес-логики

type SocialMedia int64

const (
	All SocialMedia = iota
	Telegram
	Instagram
	Youtube
)

func (sm SocialMedia) String() string {
	switch sm {
	case Telegram:
		return "tg"
	case Instagram:
		return "inst"
	case Youtube:
		return "youtube"
	default:
		return ""
	}
}

func NewSocialMedia(sm string) SocialMedia {
	switch sm {
	case "tg":
		return Telegram
	case "inst":
		return Instagram
	case "youtube":
		return Youtube
	}

	return All
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
	// тип сортировки
	Sort string
}

type CreateDoctorRequest struct {
	// username
	Telegram  string
	Instagram string
	YouTube   string
}

type UpdateDoctorRequest struct {
	// username telegram
	Telegram string
	// username telegram
	Instagram string
}
