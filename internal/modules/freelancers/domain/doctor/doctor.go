package doctor

// Doctor Доменная модель врача
type Doctor struct {
	medblogersID MedblogersID // уникальный ID в системе MEDBLOGERS

	name string // Имя
	slug string // URL

	cityID                    int64 // Основной город
	cityName                  string
	additionalCitiesIDs       []int64 // Доп.Города
	specialityID              int64   // Основная специальность
	specialityName            string
	additionalSpecialitiesIDs []int64 // Доп.Специальности

	s3Image S3Key // ссылка на S3
}

func New(options ...Option) *Doctor {
	d := &Doctor{}
	for _, option := range options {
		option(d)
	}
	return d
}
