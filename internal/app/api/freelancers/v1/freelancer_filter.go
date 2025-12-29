package v1

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/filter_freelancers/dto"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"

	"github.com/samber/lo"
)

func (i *Implementation) Filter(ctx context.Context, request *desc.FilterRequest) (*desc.FilterResponse, error) {
	filter := i.requestToFilterDTO(request)

	filterResultDomain, err := i.freelancers.Actions.FilterFreelancers.Do(ctx, filter)
	if err != nil {
		return nil, err
	}

	filterDTO := i.newFilterResponse(filterResultDomain)
	return filterDTO, nil
}

func (i *Implementation) requestToFilterDTO(req *desc.FilterRequest) freelancer.Filter {
	page := req.Page
	if page <= 0 {
		page = 1
	}

	return freelancer.Filter{
		Page:           page,
		Cities:         req.Cities,
		Specialities:   req.Specialities,
		SocialNetworks: req.Societies,
		PriceCategory:  req.PriceCategory,
	}
}

func (i *Implementation) newFilterResponse(freelancers []dto.Freelancer) *desc.FilterResponse {
	freelancersResponse := make([]*desc.FilterResponse_FreelancerItem, 0, len(freelancers))
	for _, item := range freelancers {
		freelancersResponse = append(freelancersResponse, &desc.FilterResponse_FreelancerItem{
			Name: item.Name,
			Slug: item.Slug,
			Speciality: lo.Map(item.Specialities, func(item dto.Speciality, _ int) *desc.SpecialityItem {
				return &desc.SpecialityItem{
					Id:   item.ID,
					Name: item.Name,
				}
			}),
			City: lo.Map(item.Cities, func(item dto.City, index int) *desc.CityItem {
				return &desc.CityItem{
					Id:   item.ID,
					Name: item.Name,
				}
			}),
			Image:         item.Image,
			PriceCategory: item.PriceCategory,
			SocialNetworks: lo.Map(item.Networks, func(item dto.NetworkItem, index int) *desc.FilterResponse_FreelancerItem_SocialNetworkItem {
				return &desc.FilterResponse_FreelancerItem_SocialNetworkItem{
					Id:   item.ID,
					Name: item.Name,
					Slug: item.Slug,
				}
			}),
			AgencyRepresentative: item.AgencyRepresentative,
		})
	}

	return &desc.FilterResponse{
		Freelancers: freelancersResponse,
	}
}
