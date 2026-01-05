package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/city"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/city/v1"
)

func (i *Implementation) SearchCities(ctx context.Context, req *desc.SearchCitiesRequest) (resp *desc.SearchCitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancers/cities/search", func(ctx context.Context) error {
		res, err := i.admin.Actions.FreelancerModule.CityAgg.SearchCities.Do(ctx, req.GetQuery())
		if err != nil {
			return err
		}

		resp = &desc.SearchCitiesResponse{
			Cities: lo.Map(res, func(item *city.City, _ int) *desc.SearchCitiesResponse_CityItem {
				return &desc.SearchCitiesResponse_CityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}

		return nil
	})
}
