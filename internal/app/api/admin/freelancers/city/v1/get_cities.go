package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/city/v1"
)

func (i *Implementation) GetCities(ctx context.Context, req *desc.GetCitiesRequest) (resp *desc.GetCitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancers/cities", func(ctx context.Context) error {
		res, err := i.admin.Actions.FreelancerModule.CityAgg.GetCities.Do(ctx)
		if err != nil {
			return err
		}

		resp = &desc.GetCitiesResponse{
			Cities: lo.Map(res, func(item *city.City, _ int) *desc.GetCitiesResponse_CityItem {
				return &desc.GetCitiesResponse_CityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}

		return nil
	})
}
