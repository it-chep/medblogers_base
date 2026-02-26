package doctor

import (
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"strings"
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

// GetInstURL .
func (d *Doctor) GetInstURL() string {
	return d.instURL
}

// GetProdoctorovURL .
func (d *Doctor) GetProdoctorovURL() string {
	return d.tiktokURL
}

// GetDzenURL .
func (d *Doctor) GetDzenURL() string {
	return d.dzenURL
}

// GetYoutubeURL .
func (d *Doctor) GetYoutubeURL() string {
	return d.youtubeURL
}

// GetTiktokURL .
func (d *Doctor) GetTiktokURL() string {
	return d.tiktokURL
}

// GetVkURL .
func (d *Doctor) GetVkURL() string {
	return d.vkURL
}

// GetTgURL .
func (d *Doctor) GetTgURL() string {
	if d.tgURL == "" {
		return ""
	}

	// Удаляем @ в начале, если есть
	tgURL := strings.TrimPrefix(d.tgURL, "@")

	// Если URL не содержит http/https, формируем полный URL
	if !strings.HasPrefix(tgURL, "http") {
		tgURL = "https://t.me/" + tgURL
	}

	return tgURL
}

// GetTgChannelURL возвращает полноценную ссылку на тг
func (d *Doctor) GetTgChannelURL() string {
	if d.tgChannelURL == "" {
		return ""
	}

	// Удаляем @ в начале, если есть
	tgURL := strings.TrimPrefix(d.tgChannelURL, "@")

	// Если URL не содержит http/https, формируем полный URL
	if !strings.HasPrefix(tgURL, "http") {
		tgURL = "https://t.me/" + tgURL
	}

	return tgURL
}

// GetTgChannelUsername .
func (d *Doctor) GetTgChannelUsername() string {
	return d.tgChannelURL
}

// GetSiteLink .
func (d *Doctor) GetSiteLink() string {
	return d.siteLink
}

// GetMedicalDirection .
func (d *Doctor) GetMedicalDirection() string {
	return d.medicalDirection
}

// GetMainBlogTheme .
func (d *Doctor) GetMainBlogTheme() string {
	return d.mainBlogTheme
}

// GetAdditionalCitiesIDs .
func (d *Doctor) GetAdditionalCitiesIDs() []int64 {
	return d.additionalCitiesIDs
}

// GetAdditionalSpecialitiesIDs .
func (d *Doctor) GetAdditionalSpecialitiesIDs() []int64 {
	return d.additionalSpecialitiesIDs
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

// GetIsKFDoctor получение инфы что врач из Клиники Фомина
func (d *Doctor) GetIsKFDoctor() bool {
	return d.isKFDoctor
}

// GetIsVip получение инфы об активности випки
func (d *Doctor) GetIsVip() bool {
	return d.isVip
}
