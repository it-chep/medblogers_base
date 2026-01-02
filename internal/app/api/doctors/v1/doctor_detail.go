package v1

import (
	"context"
	indto "medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"

	"github.com/samber/lo"
)

// GetDoctor - /api/v1/doctors/{doctor_id} [GET]
func (i *Implementation) GetDoctor(ctx context.Context, req *desc.GetDoctorRequest) (*desc.GetDoctorResponse, error) {
	doctorDomain, err := i.doctors.Actions.DoctorDetail.Do(ctx, req.DoctorSlug)
	if err != nil {
		return nil, err
	}

	return i.newDoctorDetailResponse(doctorDomain), nil
}

func (i *Implementation) newDoctorDetailResponse(doctorDomain *indto.DoctorDTO) *desc.GetDoctorResponse {
	if doctorDomain.InstSubsCount == "0" {
		doctorDomain.InstSubsCount = ""
	}

	if doctorDomain.TgSubsCount == "0" {
		doctorDomain.TgSubsCount = ""
	}

	if doctorDomain.YouTubeSubsCount == "0" {
		doctorDomain.YouTubeSubsCount = ""
	}

	if doctorDomain.VkSubsCount == "0" {
		doctorDomain.VkSubsCount = ""
	}

	return &desc.GetDoctorResponse{
		Name: doctorDomain.Name,
		Slug: doctorDomain.Slug,

		InstUrl:      doctorDomain.InstURL,
		VkUrl:        doctorDomain.VkURL,
		DzenUrl:      doctorDomain.DzenURL,
		YoutubeUrl:   doctorDomain.YoutubeURL,
		TgUrl:        doctorDomain.TgURL,
		TgChannelUrl: doctorDomain.TgChannelURL,
		TiktokUrl:    doctorDomain.TiktokURL,
		SiteLink:     doctorDomain.SiteLink,

		TgLastUpdatedDate: doctorDomain.TgLastUpdatedDate,
		TgSubsCountText:   doctorDomain.TgSubsCountText,
		TgSubsCount:       doctorDomain.TgSubsCount,

		InstSubsCount:       doctorDomain.InstSubsCount,
		InstSubsCountText:   doctorDomain.InstSubsCountText,
		InstLastUpdatedDate: doctorDomain.InstLastUpdatedDate,

		YoutubeSubsCountText:   doctorDomain.YouTubeSubsCountText,
		YoutubeSubsCount:       doctorDomain.YouTubeSubsCount,
		YoutubeLastUpdatedDate: doctorDomain.YouTubeLastUpdatedDate,

		VkSubsCountText:   doctorDomain.VkSubsCountText,
		VkSubsCount:       doctorDomain.VkSubsCount,
		VkLastUpdatedDate: doctorDomain.VkLastUpdatedDate,

		Cities: lo.Map(doctorDomain.Cities, func(item indto.CityItem, _ int) *desc.GetDoctorResponse_CityItem {
			return &desc.GetDoctorResponse_CityItem{
				Id:   item.ID,
				Name: item.Name,
			}
		}),

		Specialities: lo.Map(doctorDomain.Specialities, func(item indto.SpecialityItem, _ int) *desc.GetDoctorResponse_SpecialityItem {
			return &desc.GetDoctorResponse_SpecialityItem{
				Id:   item.ID,
				Name: item.Name,
			}
		}),

		MainCity: &desc.GetDoctorResponse_CityItem{
			Id:   doctorDomain.MainCity.ID,
			Name: doctorDomain.MainCity.Name,
		},

		MainSpeciality: &desc.GetDoctorResponse_SpecialityItem{
			Id:   doctorDomain.MainSpeciality.ID,
			Name: doctorDomain.MainSpeciality.Name,
		},

		MainBlogTheme: doctorDomain.MainBlogTheme,
		Image:         doctorDomain.Image,
		IsKfDoctor:    doctorDomain.IsKFDoctor,
	}
}
