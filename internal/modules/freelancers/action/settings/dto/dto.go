package dto

import (
	"github.com/samber/lo"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
)

type City struct {
	ID               int64
	Name             string
	FreelancersCount int64
}

type Speciality struct {
	ID               int64
	Name             string
	FreelancersCount int64
}

type Society struct {
	ID               int64
	Name             string
	FreelancersCount int64
}

type PriceCategory struct {
	ID               int64
	Name             string
	FreelancersCount int64
}

type Settings struct {
	Cities          []City
	Specialities    []Speciality
	SocialNetworks  []Society
	PriceCategories []PriceCategory
}

func NewSettings(cities []*city.City, specialities []*speciality.Speciality, categories []PriceCategory, networks []*social_network.SocialNetwork) *Settings {
	return &Settings{
		Cities: lo.Map(cities, func(cityItem *city.City, _ int) City {
			return City{
				ID:               cityItem.ID(),
				Name:             cityItem.Name(),
				FreelancersCount: cityItem.FreelancersCount(),
			}
		}),
		Specialities: lo.Map(specialities, func(specialityItem *speciality.Speciality, _ int) Speciality {
			return Speciality{
				ID:               specialityItem.ID(),
				Name:             specialityItem.Name(),
				FreelancersCount: specialityItem.FreelancersCount(),
			}
		}),
		SocialNetworks: lo.Map(networks, func(item *social_network.SocialNetwork, index int) Society {
			return Society{
				ID:               item.ID(),
				Name:             item.Name(),
				FreelancersCount: item.FreelancersCount(),
			}
		}),
		PriceCategories: categories,
	}
}
