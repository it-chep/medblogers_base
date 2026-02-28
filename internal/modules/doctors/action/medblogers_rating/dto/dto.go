package dto

type RatingItem struct {
	DoctorID       int64
	Slug           string
	Name           string
	S3Image        string
	Image          string
	CityID         int64
	CityName       string
	SpecialityID   int64
	SpecialityName string
	MBCCoins       int64
}

type MBC struct {
	DoctorID int64
	MBCCount int64
}
