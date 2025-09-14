package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/freelancers/domain/city"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
)

func (i *Implementation) GetCities(ctx context.Context, _ *desc.GetCitiesRequest) (*desc.CitiesResponse, error) {
	cities, err := i.freelancers.Actions.GetAllCities.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.CitiesResponse{
		Cities: lo.Map(cities, func(item *city.City, _ int) *desc.CitiesResponse_CityItem {
			return &desc.CitiesResponse_CityItem{
				CityId:   item.ID(),
				CityName: item.Name(),
			}
		}),
	}, nil
}
