package doctor

import (
	"time"

	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

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

	s3Image   S3Key     // ссылка на S3
	birthDate time.Time // дата рождения

	cooperationType CooperationType // Тип размещения

	isKFDoctor           bool // доктор из клиники фомина
	createdAt            time.Time
	marketingPreferences string
}

func New(options ...Option) *Doctor {
	d := &Doctor{}
	for _, option := range options {
		option(d)
	}
	return d
}
