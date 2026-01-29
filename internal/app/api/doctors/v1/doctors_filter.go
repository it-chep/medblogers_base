package v1

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"

	"github.com/samber/lo"
)

// Filter - /api/v1/doctors/filter [GET]
func (i *Implementation) Filter(ctx context.Context, req *desc.FilterRequest) (*desc.FilterResponse, error) {
	filter := i.requestToFilterDTO(req)

	filterResultDomain, err := i.doctors.Actions.DoctorsFilter.Do(ctx, filter)
	if err != nil {
		return nil, err
	}

	filterDTO := i.newFilterResponse(filterResultDomain)
	return filterDTO, nil
}

func (i *Implementation) requestToFilterDTO(req *desc.FilterRequest) dto.Filter {
	page := req.Page
	if page <= 0 {
		page = 1
	}

	maxSubscribers := req.MaxSubscribers
	if maxSubscribers <= 0 {
		maxSubscribers = 5_000_000
	}

	minSubscribers := req.MinSubscribers
	if minSubscribers <= 0 {
		minSubscribers = 100
	}

	return dto.Filter{
		MaxSubscribers: maxSubscribers,
		MinSubscribers: minSubscribers,
		Page:           page,
		Cities:         req.Cities,
		Specialities:   req.Specialities,
		SocialMedia:    req.SocialMedia,
		Sort:           dto.Sort(req.Sort),
	}
}

func (i *Implementation) newFilterResponse(filterDomain dto.Response) *desc.FilterResponse {
	doctorsResponse := make([]*desc.FilterResponse_DoctorItem, 0, len(filterDomain.Doctors))
	for _, item := range filterDomain.Doctors {
		if item.InstSubsCount == "0" {
			item.InstSubsCount = ""
		}

		if item.TgSubsCount == "0" {
			item.TgSubsCount = ""
		}

		if item.YouTubeSubsCount == "0" {
			item.YouTubeSubsCount = ""
		}

		if item.VkSubsCount == "0" {
			item.VkSubsCount = ""
		}

		doctorsResponse = append(doctorsResponse, &desc.FilterResponse_DoctorItem{
			Id:   item.ID,
			Name: item.Name,
			Slug: item.Slug,

			InstLink:          item.InstLink,
			InstSubsCount:     item.InstSubsCount,
			InstSubsCountText: item.InstSubsCountText,

			TgLink:          item.TgLink,
			TgSubsCount:     item.TgSubsCount,
			TgSubsCountText: item.TgSubsCountText,

			YoutubeLink:          item.YouTubeLink,
			YoutubeSubsCount:     item.YouTubeSubsCount,
			YoutubeSubsCountText: item.YouTubeSubsCountText,

			VkLink:          item.VkLink,
			VkSubsCount:     item.VkSubsCount,
			VkSubsCountText: item.VkSubsCountText,

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
			Image:      item.Image,
			IsKfDoctor: item.IsKFDoctor,
		})
	}

	return &desc.FilterResponse{
		Doctors:          doctorsResponse,
		SubscribersCount: filterDomain.SubscribersCount,
	}
}
