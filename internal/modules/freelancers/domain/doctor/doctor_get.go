package doctor

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

// GetAdditionalCitiesIDs .
func (d *Doctor) GetAdditionalCitiesIDs() []int64 {
	return d.additionalCitiesIDs
}

// GetAdditionalSpecialitiesIDs .
func (d *Doctor) GetAdditionalSpecialitiesIDs() []int64 {
	return d.additionalSpecialitiesIDs
}

// GetMainSpecialityID основная специальность
func (d *Doctor) GetMainSpecialityID() int64 {
	return d.specialityID
}

// GetMainCityID основной город
func (d *Doctor) GetMainCityID() int64 {
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
