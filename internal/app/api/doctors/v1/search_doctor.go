package v1

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/search_doctor/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Search - /api/v1/doctors/search [GET]
func (i *Implementation) Search(ctx context.Context, req *desc.SearchRequest) (*desc.SearchResponse, error) {
	if req.Query == "" {
		return nil, status.Error(codes.InvalidArgument, "query is required")
	}
	searchResultDomain, err := i.doctors.Actions.SearchDoctor.Do(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	return i.newSearchResponse(searchResultDomain), nil
}

func (i *Implementation) newSearchResponse(searchResultDomain dto.SearchDTO) *desc.SearchResponse {
	return &desc.SearchResponse{
		Doctors: lo.Map(searchResultDomain.Doctors, func(item dto.DoctorItem, _ int) *desc.SearchResponse_DoctorItem {
			return &desc.SearchResponse_DoctorItem{
				Id:             item.ID,
				Name:           item.Name,
				Slug:           item.Slug,
				CityName:       item.CityName,
				SpecialityName: item.SpecialityName,
				Image:          item.S3Image,
			}
		}),
		Cities: lo.Map(searchResultDomain.Cities, func(cityItem dto.CityItem, _ int) *desc.SearchResponse_CityItem {
			return &desc.SearchResponse_CityItem{
				Id:           cityItem.ID,
				Name:         cityItem.Name,
				DoctorsCount: cityItem.DoctorsCount,
			}
		}),
		Specialities: lo.Map(searchResultDomain.Specialities, func(specialityItem dto.SpecialityItem, _ int) *desc.SearchResponse_SpecialityItem {
			return &desc.SearchResponse_SpecialityItem{
				Id:           specialityItem.ID,
				Name:         specialityItem.Name,
				DoctorsCount: specialityItem.DoctorsCount,
			}
		}),
	}
}
