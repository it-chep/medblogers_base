package doctor

import (
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

// Doctor Доменная модель врача
type Doctor struct {
	medblogersID MedblogersID // уникальный ID в системе MEDBLOGERS
	isActive     bool         // Признак активности врача

	name string // Имя
	slug string // URL

	cityID         city.CityID // Основной город
	cityName       string
	specialityID   speciality.SpecialityID // Основная специальность
	specialityName string

	s3Image S3Key // ссылка на S3
}

func New(options ...Option) *Doctor {
	d := &Doctor{}
	for _, option := range options {
		option(d)
	}
	return d
}
