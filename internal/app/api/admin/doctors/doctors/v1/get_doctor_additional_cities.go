package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) GetDoctorAdditionalCities(ctx context.Context, req *desc.GetDoctorAdditionalCitiesRequest) (resp *desc.GetDoctorAdditionalCitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/additional_cities", func(ctx context.Context) error {

		cities, err := i.admin.Actions.DoctorModule.DoctorAgg.GetDoctorAdditionalCities.Do(ctx, req.GetDoctorId())
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorAdditionalCitiesResponse{
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
