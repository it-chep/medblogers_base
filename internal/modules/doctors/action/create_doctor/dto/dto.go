package dto

import (
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"time"
)

type CreateDoctorRequest struct {
	ID       doctor.MedblogersID
	Slug     string
	FullName string

	AgreePolicy bool

	CityID       int64
	SpecialityID int64

	Email      string
	LastName   string
	FirstName  string
	MiddleName string

	BirthDateTime   time.Time // поле заполняется при валидации
	BirthDateString string

	InstagramUsername string
	VKUsername        string
	TelegramUsername  string
	DzenUsername      string
	YoutubeUsername   string
	TelegramChannel   string
	TikTokURL         string

	MainBlogTheme string
	SiteLink      string

	AdditionalCities      []int64
	AdditionalSpecialties []int64
}

type ValidationError struct {
	Code  int
	Text  string
	Field string
}

func (e ValidationError) Error() string {
	return e.Text
}
