package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/speciality"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancerAdditionalSpecialities(ctx context.Context, req *desc.GetFreelancerAdditionalSpecialitiesRequest) (resp *desc.GetFreelancerAdditionalSpecialitiesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/additional_specialities", func(ctx context.Context) error {
		specialities, err := i.admin.Actions.FreelancerModule.FreelancerAgg.GetAdditionalSpecialities.Do(ctx, req.GetFreelancerId())
		if err != nil {
			return err
		}

		resp = &desc.GetFreelancerAdditionalSpecialitiesResponse{
			AdditionalSpecialities: lo.Map(specialities, func(item *speciality.Speciality, index int) *desc.SpecialityItem {
				return &desc.SpecialityItem{
					Id:   int64(item.ID()),
					Name: item.Name(),
				}
			}),
		}
		return nil
	})
}
