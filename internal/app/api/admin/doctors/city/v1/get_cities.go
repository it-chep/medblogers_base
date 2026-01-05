package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/city/v1"
)

func (i *Implementation) GetCities(ctx context.Context, req *desc.GetDoctorsCitiesRequest) (resp *desc.GetDoctorsCitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors/cities", func(ctx context.Context) error {
		res, err := i.admin.Actions.DoctorModule.CityAgg.GetCities.Do(ctx)
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorsCitiesResponse{
			Cities: lo.Map(res, func(item *city.City, _ int) *desc.GetDoctorsCitiesResponse_CityItem {
				return &desc.GetDoctorsCitiesResponse_CityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}

		return nil
	})
}
