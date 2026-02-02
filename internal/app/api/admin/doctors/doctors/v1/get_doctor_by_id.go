package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) GetDoctorByID(ctx context.Context, req *desc.GetDoctorByIDRequest) (resp *desc.GetDoctorByIDResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}", func(ctx context.Context) error {
		docDTO, err := i.admin.Actions.DoctorModule.DoctorAgg.GetDoctorByID.Do(ctx, req.GetDoctorId())
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorByIDResponse{
			Id:   docDTO.ID,
			Name: docDTO.Name,
			Slug: docDTO.Slug,

			// CooperationType:
			InstUrl:      docDTO.InstURL,
			TgUrl:        docDTO.TgURL,
			VkUrl:        docDTO.VkURL,
			YoutubeUrl:   docDTO.YoutubeURL,
			TiktokUrl:    docDTO.TikTokURL,
			DzenUrl:      docDTO.DzenURL,
			TgChannelUrl: docDTO.TgChannelURL,
			SiteLink:     docDTO.SiteLink,

			//AdditionalSpecialities: lo.Map(docDTO.AdditionalSpecialities, func(item dto.Speciality, _ int) *desc.SpecialityItem {
			//	return &desc.SpecialityItem{
			//		Id:   item.ID,
			//		Name: item.Name,
			//	}
			//}),
			//AdditionalCities: lo.Map(docDTO.AdditionalCities, func(item dto.City, _ int) *desc.CityItem {
			//	return &desc.CityItem{
			//		Id:   item.ID,
			//		Name: item.Name,
			//	}
			//}),

			MainCity: &desc.CityItem{
				Id:   docDTO.MainCity.ID,
				Name: docDTO.MainCity.Name,
			},
			MainSpeciality: &desc.SpecialityItem{
				Id:   docDTO.MainSpeciality.ID,
				Name: docDTO.MainSpeciality.Name,
			},
			MarketingPreferences: docDTO.MarketingPreferences,
			MainBlogTheme:        docDTO.MainBlogTheme,
			MedicalDirections:    docDTO.MedicalDirections,

			SubscribersInfo: lo.Map(docDTO.SubscribersInfo, func(item dto.Subscribers, _ int) *desc.GetDoctorByIDResponse_SubscribersItem {
				return &desc.GetDoctorByIDResponse_SubscribersItem{
					Key:             item.Key,
					SubsCount:       item.SubsCount,
					SubsCountText:   item.SubsCountText,
					LastUpdatedDate: item.LastUpdatedDate,
				}
			}),

			Image:      docDTO.Image,
			IsKfDoctor: docDTO.IsKfDoctor,
			IsActive:   docDTO.IsActive,
		}
		return nil
	})
}
