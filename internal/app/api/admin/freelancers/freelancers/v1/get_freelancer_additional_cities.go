package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancerAdditionalCities(ctx context.Context, req *desc.GetFreelancerAdditionalCitiesRequest) (resp *desc.GetFreelancerAdditionalCitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/additional_cities", func(ctx context.Context) error {

		cities, err := i.admin.Actions.FreelancerModule.FreelancerAgg.GetAdditionalCities.Do(ctx, req.GetFreelancerId())
		if err != nil {
			return err
		}

		resp = &desc.GetFreelancerAdditionalCitiesResponse{
			AdditionalCities: lo.Map(cities, func(item *city.City, index int) *desc.CityItem {
				return &desc.CityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}
		return nil
	})
}
