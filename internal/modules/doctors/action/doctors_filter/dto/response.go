package dto

type Doctor struct {
	ID   int64
	Name string
	Slug string

	InstLink          string
	InstSubsCount     string
	InstSubsCountText string

	TgLink          string
	TgSubsCount     string
	TgSubsCountText string

	Specialities []Speciality // Строка из основной и дополнительных специальностей
	Cities       []City       // Строка из основного и дополнительных городов

	Image string

	MainCityID       int64
	MainSpecialityID int64

	S3Key string

	IsKFDoctor bool
}

type City struct {
	ID   int64
	Name string
}

type Speciality struct {
	ID   int64
	Name string
}

type Response struct {
	Doctors          []Doctor
	Pages            int64
	SubscribersCount string
}
