package v1

import (
	"context"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
)

func (i *Implementation) Search(ctx context.Context, request *desc.SearchRequest) (*desc.SearchResponse, error) {
	if request.Query == "" {
		return nil, status.Error(codes.InvalidArgument, "query is required")
	}
	searchResultDomain, err := i.freelancers.Actions.SearchFreelancers.Do(ctx, request.Query)
	if err != nil {
		return nil, err
	}
	return i.newSearchResponse(searchResultDomain), nil
}

func (i *Implementation) newSearchResponse(searchResultDomain dto.SearchDTO) *desc.SearchResponse {
	return &desc.SearchResponse{
		Freelancers: lo.Map(searchResultDomain.Freelancers, func(item dto.FreelancerItem, _ int) *desc.SearchResponse_FreelancerItem {
			return &desc.SearchResponse_FreelancerItem{
				Id:                    item.ID,
				Name:                  item.Name,
				Slug:                  item.Slug,
				CityName:              item.CityName,
				SpecialityName:        item.SpecialityName,
				Image:                 item.S3Image,
				ExperienceWithDoctors: item.ExperienceWithDoctors,
				PriceCategory:         item.PriceCategory,
				SocialNetworks:        nil,
			}
		}),
		Cities: lo.Map(searchResultDomain.Cities, func(cityItem dto.CityItem, _ int) *desc.SearchResponse_CityItem {
			return &desc.SearchResponse_CityItem{
				Id:               cityItem.ID,
				Name:             cityItem.Name,
				FreelancersCount: cityItem.FreelancersCount,
			}
		}),
		Specialities: lo.Map(searchResultDomain.Specialities, func(specialityItem dto.SpecialityItem, _ int) *desc.SearchResponse_SpecialityItem {
			return &desc.SearchResponse_SpecialityItem{
				Id:               specialityItem.ID,
				Name:             specialityItem.Name,
				FreelancersCount: specialityItem.FreelancersCount,
			}
		}),
	}
}
