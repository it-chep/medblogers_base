package v1

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/city"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"

	"github.com/samber/lo"
)

// GetCities - /api/v1/cities_list [GET]
func (i *Implementation) GetCities(ctx context.Context, _ *desc.GetCitiesRequest) (*desc.CitiesResponse, error) {
	cities, err := i.doctors.Actions.AllCities.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.CitiesResponse{
		Cities: lo.Map(cities, func(item *city.City, _ int) *desc.CitiesResponse_CityItem {
			return &desc.CitiesResponse_CityItem{
				CityId:   int64(item.ID()),
				CityName: item.Name(),
			}
		}),
	}, nil
}
