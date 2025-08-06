package dto

import (
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"

	"github.com/samber/lo"
)

type FilterItem struct {
	Name string
	Slug string
}

type CityItem struct {
	ID           int64
	Name         string
	DoctorsCount int64
}

type SpecialityItem struct {
	ID           int64
	Name         string
	DoctorsCount int64
}

type Settings struct {
	FilterInfo      []FilterItem
	Cities          []CityItem
	Specialities    []SpecialityItem
	NewDoctorBanner bool
}

func NewSettings(cities []*city.City, specialities []*speciality.Speciality, filters []indto.FilterInfoResponse) *Settings {
	return &Settings{
		FilterInfo: lo.Map(filters, func(item indto.FilterInfoResponse, index int) FilterItem {
			return FilterItem{
				Name: item.Name,
				Slug: item.Slug,
			}
		}),
		Cities: lo.Map(cities, func(cityItem *city.City, _ int) CityItem {
			return CityItem{
				ID:           int64(cityItem.ID()),
				Name:         cityItem.Name(),
				DoctorsCount: cityItem.DoctorsCount(),
			}
		}),
		Specialities: lo.Map(specialities, func(specialityItem *speciality.Speciality, _ int) SpecialityItem {
			return SpecialityItem{
				ID:           int64(specialityItem.ID()),
				Name:         specialityItem.Name(),
				DoctorsCount: specialityItem.DoctorsCount(),
			}
		}),
		NewDoctorBanner: true,
	}
}
