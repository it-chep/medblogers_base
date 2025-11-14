package v1

import (
	"context"
	indto "medblogers_base/internal/modules/freelancers/action/settings/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"

	"github.com/samber/lo"
)

func (i *Implementation) GetSettings(ctx context.Context, _ *desc.GetSettingsRequest) (*desc.GetSettingsResponse, error) {
	settingsDomain, err := i.freelancers.Actions.GetSettings.Do(ctx)
	if err != nil {
		return nil, err
	}

	settingsResponse := i.newSettingsResponse(settingsDomain)
	return settingsResponse, nil
}

func (i *Implementation) newSettingsResponse(settingsDomain *indto.Settings) *desc.GetSettingsResponse {
	return &desc.GetSettingsResponse{
		Cities: lo.Map(settingsDomain.Cities, func(cityItem indto.City, _ int) *desc.GetSettingsResponse_CityItem {
			return &desc.GetSettingsResponse_CityItem{
				Id:               cityItem.ID,
				Name:             cityItem.Name,
				FreelancersCount: cityItem.FreelancersCount,
			}
		}),
		Specialities: lo.Map(settingsDomain.Specialities, func(specialityItem indto.Speciality, _ int) *desc.GetSettingsResponse_SpecialityItem {
			return &desc.GetSettingsResponse_SpecialityItem{
				Id:               specialityItem.ID,
				Name:             specialityItem.Name,
				FreelancersCount: specialityItem.FreelancersCount,
			}
		}),
		Societies: lo.Map(settingsDomain.SocialNetworks, func(item indto.Society, index int) *desc.GetSettingsResponse_SocietyItem {
			return &desc.GetSettingsResponse_SocietyItem{
				Id:               item.ID,
				Name:             item.Name,
				FreelancersCount: item.FreelancersCount,
			}
		}),
		PriceCategories: lo.Map(settingsDomain.PriceCategories, func(item indto.PriceCategory, index int) *desc.GetSettingsResponse_PriceCategoryItem {
			return &desc.GetSettingsResponse_PriceCategoryItem{
				Id:               item.ID,
				Name:             item.Name,
				FreelancersCount: item.FreelancersCount,
			}
		}),
	}
}
