package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) UpdateDoctor(ctx context.Context, req *desc.UpdateDoctorRequest) (resp *desc.UpdateDoctorResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/update", func(ctx context.Context) error {

		updateReq := getUpdateDoctorReq(req)
		err = i.admin.Actions.DoctorModule.DoctorAgg.UpdateDoctor.Do(ctx, req.GetDoctorId(), updateReq)
		if err != nil {
			return err
		}

		return nil
	})
}

func getUpdateDoctorReq(req *desc.UpdateDoctorRequest) dto.UpdateRequest {
	return dto.UpdateRequest{
		Name:                 req.GetName(),
		Slug:                 req.GetSlug(),
		InstURL:              req.GetInstUrl(),
		VkURL:                req.GetVkUrl(),
		DzenURL:              req.GetDzenUrl(),
		TgURL:                req.GetTgUrl(),
		TgChannelURL:         req.GetTgChannelUrl(),
		YouTubeURL:           req.GetYoutubeUrl(),
		TikTokURL:            req.GetTiktokUrl(),
		SiteLink:             req.GetSiteLink(),
		MainCityID:           req.GetMainCityId(),
		MainSpecialityID:     req.GetMainSpecialityId(),
		MainBlogTheme:        req.GetMainBlogTheme(),
		IsKFDoctor:           req.GetIsKfDoctor(),
		BirthDate:            req.GetBirthDate(),
		CooperationTypeID:    req.GetCooperationTypeId(),
		MedicalDirections:    req.GetMedicalDirections(),
		MarketingPreferences: req.GetMarketingPreferences(),
	}
}
