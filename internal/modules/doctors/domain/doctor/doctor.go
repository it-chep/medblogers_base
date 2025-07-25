package doctor

import (
	"time"

	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

// CooperationType - Типы размещения врачей
type CooperationType int

const (
	// UNKNOWN - Неизвестно
	UNKNOWN CooperationType = 1
	// Club6Plus - Клуб 6+ месяцев
	Club6Plus CooperationType = 2
	// AdvertisingBarter - Бартер по рекламе
	AdvertisingBarter CooperationType = 3
	// PaidForever - Платно навсегда
	PaidForever CooperationType = 4
	// Subscription - Подписка
	Subscription CooperationType = 5
	// ReadydocPardon - Помилование от READYDOC
	ReadydocPardon CooperationType = 6
)

func (c CooperationType) String() string {
	switch c {
	case Club6Plus:
		return "Клуб 6+ месяцев"
	case AdvertisingBarter:
		return "Бартер по рекламе"
	case PaidForever:
		return "Платно навсегда"
	case Subscription:
		return "Подписка"
	case ReadydocPardon:
		return "Помилование от READYDOC"
	}
	return "Неизвестно"
}

// Doctor Доменная модель врача
type Doctor struct {
	medblogersID MedblogersID // уникальный ID в системе MEDBLOGERS
	isActive     bool         // Признак активности врача

	name  string // Имя
	slug  string // URL
	email string // Email

	// соц.сети
	instURL      string // инст
	vkURL        string // вк
	dzenURL      string // дзен
	tgURL        string // тг личный
	tgChannelURL string // канал тг
	youtubeURL   string // ютуб
	tiktokURL    string // тикток
	siteLink     string // Ссылка на личный сайт

	cityID                    city.CityID // Основной город
	cityName                  string
	additionalCitiesIDs       []int64                 // Доп.Города
	specialityID              speciality.SpecialityID // Основная специальность
	specialityName            string
	additionalSpecialitiesIDs []int64 // Доп.Специальности

	medicalDirection string // Направление медицины
	mainBlogTheme    string // Тематика блога

	s3Image   string    // ссылка на S3
	birthDate time.Time // дата рождения

	cooperationType CooperationType // Тип размещения
}

func New(options ...Option) *Doctor {
	d := &Doctor{}
	for _, option := range options {
		option(d)
	}
	return d
}
