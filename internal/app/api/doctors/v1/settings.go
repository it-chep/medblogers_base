package v1

import (
	"context"
	indto "medblogers_base/internal/modules/doctors/action/settings/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"

	"github.com/samber/lo"
)

// GetSettings - /api/v1/settings [GET]
func (i *Implementation) GetSettings(ctx context.Context, _ *desc.GetSettingsRequest) (*desc.GetSettingsResponse, error) {
	settingsDomain, err := i.doctors.Actions.Settings.Do(ctx)
	if err != nil {
		return nil, err
	}

	settingsResponse := i.newSettingsResponse(settingsDomain)
	return settingsResponse, nil
}

func (i *Implementation) newSettingsResponse(settingsDomain *indto.Settings) *desc.GetSettingsResponse {
	return &desc.GetSettingsResponse{
		NewDoctorBanner: settingsDomain.NewDoctorBanner,
		FilterInfo: lo.Map(settingsDomain.FilterInfo, func(filterItem indto.FilterItem, _ int) *desc.GetSettingsResponse_FilterItem {
			return &desc.GetSettingsResponse_FilterItem{
				Name: filterItem.Name,
				Slug: filterItem.Slug,
			}
		}),
		Cities: lo.Map(settingsDomain.Cities, func(cityItem indto.CityItem, _ int) *desc.GetSettingsResponse_CityItem {
			return &desc.GetSettingsResponse_CityItem{
				Id:           cityItem.ID,
				Name:         cityItem.Name,
				DoctorsCount: cityItem.DoctorsCount,
			}
		}),
		Specialities: lo.Map(settingsDomain.Specialities, func(specialityItem indto.SpecialityItem, _ int) *desc.GetSettingsResponse_SpecialityItem {
			return &desc.GetSettingsResponse_SpecialityItem{
				Id:           specialityItem.ID,
				Name:         specialityItem.Name,
				DoctorsCount: specialityItem.DoctorsCount,
			}
		}),
	}
}
