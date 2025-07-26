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
	return dto.Filter{
		MaxSubscribers: req.MaxSubscribers,
		MinSubscribers: req.MinSubscribers,
		Page:           req.Page,
		Cities:         req.Cities,
		Specialities:   req.Specialities,
		SocialMedia:    req.SocialMedia,
	}
}

func (i *Implementation) newFilterResponse(filterDomain dto.Response) *desc.FilterResponse {
	return &desc.FilterResponse{
		Doctors: lo.Map(filterDomain.Doctors, func(item dto.Doctor, _ int) *desc.FilterResponse_DoctorItem {
			return &desc.FilterResponse_DoctorItem{
				Id:                item.ID,
				Name:              item.Name,
				Slug:              item.Slug,
				InstLink:          item.InstLink,
				InstSubsCount:     item.InstSubsCount,
				InstSubsCountText: item.InstSubsCountText,
				TgLink:            item.TgLink,
				TgSubsCount:       item.TgSubsCount,
				TgSubsCountText:   item.TgSubsCountText,
				Speciality:        item.Speciality,
				City:              item.City,
				Image:             item.Image,
			}
		}),
		CurrentPage: filterDomain.CurrentPage,
		Pages:       filterDomain.Pages,
	}
}
