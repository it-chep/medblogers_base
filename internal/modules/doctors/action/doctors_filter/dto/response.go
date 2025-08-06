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

	Speciality string // Строка из основной и дополнительных специальностей
	City       string // Строка из основного и дополнительных городов

	Image string

	MainCityID       int64
	MainSpecialityID int64
}

type Response struct {
	Doctors          []Doctor
	Pages            int64
	CurrentPage      int64
	SubscribersCount string
}
