package doctor_author

func (d *Doctor) GetName() string {
	return d.name
}

func (d *Doctor) GetSlug() string {
	return d.slug
}

func (d *Doctor) GetS3Key() string {
	return d.s3key
}

func (d *Doctor) GetSpecialityName() string {
	return d.specialityName
}
