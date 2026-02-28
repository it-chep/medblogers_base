package dto

type RatingItem struct {
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
