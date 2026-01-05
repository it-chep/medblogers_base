package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancers(ctx context.Context, req *desc.GetFreelancersRequest) (resp *desc.GetFreelancersResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancers", func(ctx context.Context) error {
		res, err := i.admin.Actions.FreelancerModule.FreelancerAgg.GetFreelancers.Do(ctx)
		if err != nil {
			return err
		}

		resp = &desc.GetFreelancersResponse{
			Freelancers: lo.Map(res, func(item *freelancer.Freelancer, index int) *desc.GetFreelancersResponse_Freelancer {
				return &desc.GetFreelancersResponse_Freelancer{
					Id:       item.GetID(),
					Name:     item.GetName(),
					Image:    item.GetS3Image(),
					IsActive: item.GetIsActive(),
					//CooperationType: &desc.GetFreelancersResponse_Freelancer_CooperationType{
					//	Id: ,
					//	Name: ,
					//},
				}
			}),
		}

		return nil
	})
}
