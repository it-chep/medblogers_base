package dto

import "medblogers_base/internal/modules/doctors/domain/doctor"

type CreateDoctorRequest struct {
	ID       doctor.MedblogersID
	Slug     string
	FullName string

	Email                 string
	LastName              string
	FirstName             string
	MiddleName            string
	BirthDate             string
	AdditionalCities      []int64
	AdditionalSpecialties []int64
	InstagramUsername     string
	VKUsername            string
	TelegramUsername      string
	DzenUsername          string
	YoutubeUsername       string
	TelegramChannel       string
	CityID                int64
	SpecialityID          int64
	MainBlogTheme         string
	SiteLink              string
	AgreePolicy           bool
}
