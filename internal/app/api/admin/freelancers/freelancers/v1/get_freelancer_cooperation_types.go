package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancerCooperationTypes(ctx context.Context, _ *desc.GetFreelancerCooperationTypesRequest) (resp *desc.GetFreelancerCooperationTypesResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancer/cooperation_types", func(ctx context.Context) error {
		coopTypes, err := i.admin.Actions.FreelancerModule.FreelancerAgg.GetCooperationTypes.Do(ctx)
		if err != nil {
			return err
		}

		resp = &desc.GetFreelancerCooperationTypesResponse{
			CooperationTypes: lo.Map(coopTypes, func(item *freelancer.CooperationType, index int) *desc.CooperationType {
				return &desc.CooperationType{
					Id:   item.ID(),
					Name: item.Name(),
				}
			}),
		}
		return nil
	})
}
