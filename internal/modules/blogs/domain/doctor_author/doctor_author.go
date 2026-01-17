package doctor_author

type Doctor struct {
	name           string
	slug           string
	s3key          string
	specialityName string
}

func NewDoctor(name, slug, s3Key, specialityName string) *Doctor {
	return &Doctor{
		name:           name,
		slug:           slug,
		s3key:          s3Key,
		specialityName: specialityName,
	}
}
