package indto

// Абстракция для бизнес-логики

type SocialMedia int64

const (
	All SocialMedia = iota
	Telegram
	Instagram
	Youtube
	Vk
)

func (sm SocialMedia) String() string {
	switch sm {
	case Telegram:
		return "tg"
	case Instagram:
		return "inst"
	case Youtube:
		return "youtube"
	case Vk:
		return "vk"
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
	case "vk":
		return Vk
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
	Vk        string
}

type UpdateDoctorRequest struct {
	// username telegram
	Telegram string
	// username telegram
	Instagram string

	IsActive bool
}

type UpdateSubscribersItem struct {
	Key       SocialMedia
	SubsCount int64
}
