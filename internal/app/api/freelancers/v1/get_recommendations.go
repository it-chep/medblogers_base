package v1

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/get_recommendations/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"

	"github.com/samber/lo"
)

func (i *Implementation) GetFreelancerRecommendations(ctx context.Context, req *desc.GetFreelancerRecommendationsRequest) (*desc.GetFreelancerRecommendationsResponse, error) {
	resp, err := i.freelancers.Actions.GetRecommendations.Do(ctx, req.GetFreelancerSlug())
	if err != nil {
		return nil, err
	}

	return &desc.GetFreelancerRecommendationsResponse{
		Doctors: lo.Map(resp.Doctors, func(item dto.Doctor, _ int) *desc.GetFreelancerRecommendationsResponse_Doctor {
			return &desc.GetFreelancerRecommendationsResponse_Doctor{
				Name:       item.Name,
				Slug:       item.Slug,
				Speciality: item.Speciality,
				City:       item.City,
				Image:      item.Image,
			}
		}),
	}, nil
}
