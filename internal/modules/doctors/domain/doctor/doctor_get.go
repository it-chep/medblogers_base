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
	return d.tgURL
}

// GetTgChannelURL .
func (d *Doctor) GetTgChannelURL() string {
	return d.tgChannelURL
}

// todo сделать "https валидацию"

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

// GetS3Key получение ключа для доступа к фото
func (d *Doctor) GetS3Key() string {
	return d.s3Image
}
