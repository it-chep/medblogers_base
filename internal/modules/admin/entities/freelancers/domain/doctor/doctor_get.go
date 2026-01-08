package doctor

import (
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
)

// GetID .
func (d *Doctor) GetID() MedblogersID {
	return d.medblogersID
}

// GetName .
func (d *Doctor) GetName() string {
	return d.name
}

// GetSlug .
func (d *Doctor) GetSlug() string {
	return d.slug
}

// GetMainSpecialityID основная специальность
func (d *Doctor) GetMainSpecialityID() speciality.SpecialityID {
	return d.specialityID
}

// GetMainCityID основной город
func (d *Doctor) GetMainCityID() city.CityID {
	return d.cityID
}

// GetMainCityName основной город
func (d *Doctor) GetMainCityName() string {
	return d.cityName
}

// GetMainSpecialityName основной город
func (d *Doctor) GetMainSpecialityName() string {
	return d.specialityName
}

// GetS3Key получение ключа для доступа к фото
func (d *Doctor) GetS3Key() S3Key {
	return d.s3Image
}

// GetIsActive получение активности доктора
func (d *Doctor) GetIsActive() bool {
	return d.isActive
}
