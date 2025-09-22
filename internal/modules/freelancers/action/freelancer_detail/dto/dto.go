package dto

import (
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
)

// CityItem .
type CityItem struct {
	ID   int64
	Name string
}

// SpecialityItem .
type SpecialityItem struct {
	ID   int64
	Name string
}

// PriceListItem .
type PriceListItem struct {
	Name  string
	Price int64
}

// SocialNetworkItem .
type SocialNetworkItem struct {
	ID   int64
	Name string
}

type FreelancerDTO struct {
	Name string
	Slug string

	TgURL                   string
	HasExperienceWithDoctor bool
	PriceCategory           int64
	PortfolioLink           string

	Cities         []CityItem          // доп Города
	Specialities   []SpecialityItem    // доп Специальности
	PriceList      []PriceListItem     // прайс-лист
	SocialNetworks []SocialNetworkItem // соц сети

	// основной город
	MainCityID int64
	MainCity   CityItem
	// основная специальность
	MainSpecialityID int64
	MainSpeciality   SpecialityItem

	// фотка
	Image string

	HasCommand        bool
	WorkingExperience string
}

// New .
func New(frlncr *freelancer.Freelancer) *FreelancerDTO {
	return &FreelancerDTO{
		Name: frlncr.GetName(),
		Slug: frlncr.GetSlug(),

		TgURL:                   frlncr.GetTgURL(),
		HasExperienceWithDoctor: frlncr.HasExperienceWithDoctor(),
		PriceCategory:           frlncr.GetPriceCategory(),
		PortfolioLink:           frlncr.GetPortfolioLink(),

		MainCityID:       frlncr.GetMainCityID(),
		MainSpecialityID: frlncr.GetMainSpecialityID(),

		HasCommand:        frlncr.GetHasCommand(),
		WorkingExperience: frlncr.GetWorkingExperience(),
	}
}
