package v1

import (
	"context"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/city"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/city/v1"
)

func (i *Implementation) SearchCities(ctx context.Context, req *desc.SearchDoctorsCitiesRequest) (resp *desc.SearchDoctorsCitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctors/search", func(ctx context.Context) error {
		if len(req.GetQuery()) < 2 {
			return errors.New("query is too short")
		}

		res, err := i.admin.Actions.DoctorModule.CityAgg.SearchCities.Do(ctx, req.GetQuery())
		if err != nil {
			return err
		}

		resp = &desc.SearchDoctorsCitiesResponse{
			Cities: lo.Map(res, func(item *city.City, _ int) *desc.SearchDoctorsCitiesResponse_CityItem {
				return &desc.SearchDoctorsCitiesResponse_CityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}

		return nil
	})
}
