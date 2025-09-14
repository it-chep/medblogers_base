package v1

import (
	"context"
	"github.com/samber/lo"
	indto "medblogers_base/internal/modules/freelancers/action/freelancer_detail/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
)

func (i *Implementation) GetFreelancer(ctx context.Context, request *desc.GetFreelancerRequest) (*desc.GetFreelancerResponse, error) {
	freelancer, err := i.freelancers.Actions.FreelancerDetail.Do(ctx, request.FreelancerSlug)
	if err != nil {
		return nil, err
	}

	return i.newDoctorDetailResponse(freelancer), nil
}

func (i *Implementation) newDoctorDetailResponse(freelancer *indto.FreelancerDTO) *desc.GetFreelancerResponse {

	return &desc.GetFreelancerResponse{
		Name: freelancer.Name,
		Slug: freelancer.Slug,

		TgUrl: freelancer.TgURL,

		Cities: lo.Map(freelancer.Cities, func(item indto.CityItem, _ int) *desc.GetFreelancerResponse_CityItem {
			return &desc.GetFreelancerResponse_CityItem{
				Id:   item.ID,
				Name: item.Name,
			}
		}),

		Specialities: lo.Map(freelancer.Specialities, func(item indto.SpecialityItem, _ int) *desc.GetFreelancerResponse_SpecialityItem {
			return &desc.GetFreelancerResponse_SpecialityItem{
				Id:   item.ID,
				Name: item.Name,
			}
		}),

		MainCity: &desc.GetFreelancerResponse_CityItem{
			Id:   freelancer.MainCity.ID,
			Name: freelancer.MainCity.Name,
		},

		MainSpeciality: &desc.GetFreelancerResponse_SpecialityItem{
			Id:   freelancer.MainSpeciality.ID,
			Name: freelancer.MainSpeciality.Name,
		},

		Image: freelancer.Image,
	}
}
