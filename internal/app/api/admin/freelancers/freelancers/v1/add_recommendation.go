package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) AddRecommendation(ctx context.Context, req *desc.AddRecommendationRequest) (resp *desc.AddRecommendationResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/add_recommendation", func(ctx context.Context) error {
		return i.admin.Actions.FreelancerModule.FreelancerAgg.AddRecommendation.Do(ctx, req.GetFreelancerId(), req.GetDoctorId())
	})
}
