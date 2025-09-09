package freelancer

type Freelancer struct {
	id int64

	isActive             bool  // Признак активности фрилансера
	experienceWithDoctor bool  // Опыт работы с врачами
	priceCategory        int64 // Ценовая категория

	name  string // Имя
	slug  string // URL
	email string // Email

	tgURL         string // тг личный
	portfolioLink string // Ссылка на портфолио

	cityID                    int64 // Основной город
	cityName                  string
	additionalCitiesIDs       []int64 // Доп.Города
	specialityID              int64   // Основная специальность
	specialityName            string
	additionalSpecialitiesIDs []int64 // Доп.Специальности
	socialNetworks            []int64 // Соцсети в которых работает фрилансер

	s3Image string // ссылка на S3
}

func New(options ...Option) *Freelancer {
	d := &Freelancer{}
	for _, option := range options {
		option(d)
	}
	return d
}
