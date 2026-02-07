package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_recommendations/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancerRecommendations(ctx context.Context, req *desc.GetFreelancerRecommendationsRequest) (resp *desc.GetFreelancerRecommendationsResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/recommendations", func(ctx context.Context) error {
		recommendations, err := i.admin.Actions.FreelancerModule.FreelancerAgg.GetRecommendations.Do(ctx, req.GetFreelancerId())
		if err != nil {
			return err
		}

		resp = &desc.GetFreelancerRecommendationsResponse{
			Recommendations: lo.Map(recommendations, func(item dto.Recommendation, index int) *desc.Recommendation {
				return &desc.Recommendation{
					DoctorName: item.DoctorName,
					DoctorId:   item.DoctorID,
				}
			}),
		}
		return nil
	})
}
