package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/speciality/v1"
)

func (i *Implementation) SearchSpeciality(ctx context.Context, req *desc.SearchSpecialityRequest) (resp *desc.SearchSpecialityResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancers/specialities/search", func(ctx context.Context) error {
		res, err := i.admin.Actions.FreelancerModule.SpecialityAgg.SearchSpecialities.Do(ctx, req.GetQuery())
		if err != nil {
			return err
		}

		resp = &desc.SearchSpecialityResponse{
			Specialities: lo.Map(res, func(item *speciality.Speciality, _ int) *desc.SearchSpecialityResponse_SpecialityItem {
				return &desc.SearchSpecialityResponse_SpecialityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}

		return nil
	})
}
