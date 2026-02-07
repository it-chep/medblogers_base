package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) SearchFreelancers(ctx context.Context, req *desc.SearchFreelancersRequest) (resp *desc.SearchFreelancersResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancers/search", func(ctx context.Context) error {
		res, err := i.admin.Actions.FreelancerModule.FreelancerAgg.SearchFreelancers.Do(ctx, req.GetQuery())
		if err != nil {
			return err
		}

		resp = &desc.SearchFreelancersResponse{
			Freelancers: lo.Map(res, func(item *freelancer.Freelancer, index int) *desc.SearchFreelancersResponse_Freelancer {
				return &desc.SearchFreelancersResponse_Freelancer{
					Id:       item.GetID(),
					Name:     item.GetName(),
					Image:    item.GetS3Image(),
					IsActive: item.GetIsActive(),
					//CooperationType: &desc.SearchFreelancersResponse_Freelancer_CooperationType{
					//	Id: ,
					//	Name: ,
					//},
				}
			}),
		}

		return nil
	})
}
