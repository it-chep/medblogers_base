package dto

import (
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"time"
)

type Recommendation struct {
	DoctorID   int64
	DoctorName string
}

type City struct {
	ID   int64
	Name string
}

type Speciality struct {
	ID   int64
	Name string
}

type Network struct {
	ID   int64
	Name string
}

type PriceList struct {
	ID     int64
	Name   string
	Amount string
}

type FreelancerDTO struct {
	ID int64

	IsActive             bool      // Признак активности фрилансера
	AgencyRepresentative bool      // Есть своя команда
	PriceCategory        int64     // Ценовая категория
	StartWorking         time.Time // Примерная дата начала работы

	Name  string // Имя
	Slug  string // URL
	Email string // Email

	TgURL         string // тг личный
	PortfolioLink string // Ссылка на портфолио

	City                   City         // Основной город
	AdditionalCities       []City       // Доп.Города
	Speciality             Speciality   // Основная специальность
	AdditionalSpecialities []Speciality // Доп.Специальности
	SocialNetworks         []Network    // Соцсети в которых работает фрилансер

	S3Key   string // ключик в базе
	S3Image string // ссылка на S3

	CooperationType int64

	PriceList       []PriceList
	Recommendations []Recommendation
}

func New(frlncr *freelancer.Freelancer) *FreelancerDTO {
	return &FreelancerDTO{
		ID:                   frlncr.GetID(),
		IsActive:             frlncr.IsActive(),
		AgencyRepresentative: frlncr.IsAgencyRepresentative(),
		PriceCategory:        frlncr.GetPriceCategory(),
		StartWorking:         frlncr.GetStartWorking(),
		Name:                 frlncr.GetName(),
		Slug:                 frlncr.GetSlug(),
		Email:                frlncr.GetEmail(),
		TgURL:                frlncr.GetTg(),
		PortfolioLink:        frlncr.GetPortfolioLink(),
		City:                 City{ID: frlncr.GetCityID()},
		Speciality:           Speciality{ID: frlncr.GetSpecialityID()},
		S3Key:                frlncr.GetS3Image(),
		CooperationType:      frlncr.GetCooperationType(),
	}
}
