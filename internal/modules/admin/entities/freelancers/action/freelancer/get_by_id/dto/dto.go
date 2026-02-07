package dto

import (
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"time"
)

type City struct {
	ID   int64
	Name string
}

type Speciality struct {
	ID   int64
	Name string
}

type CooperationType struct {
	ID   int64
	Name string
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

	City       City       // Основной город
	Speciality Speciality // Основная специальность

	S3Key   string // ключик в базе
	S3Image string // ссылка на S3

	CooperationType CooperationType
	CreatedAt       time.Time
}

func New(frlncr *freelancer.Freelancer) *FreelancerDTO {
	cooperationType := frlncr.GetCooperationType()
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
		CooperationType: CooperationType{
			ID:   cooperationType.ID(),
			Name: cooperationType.Name(),
		},
		CreatedAt: frlncr.GetCreatedAt(),
	}
}
