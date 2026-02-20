package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
	"time"
)

func (i *Implementation) GetDoctorVipInfo(ctx context.Context, req *desc.GetDoctorVipInfoRequest) (resp *desc.GetDoctorVipInfoResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{doctor_id}/vip_info", func(ctx context.Context) error {
		vipInfo, err := i.admin.Actions.DoctorModule.DoctorAgg.GetDoctorVipInfo.Do(ctx, req.GetDoctorId())
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorVipInfoResponse{
			CanBarter:            vipInfo.GetCanBarter(),
			CanBuyAdvertising:    vipInfo.GetCanBuyAdvertising(),
			CanSellAdvertising:   vipInfo.GetCanSellAdvertising(),
			AdvertisingPriceFrom: vipInfo.GetAdvertisingPriceFrom(),
			BlogInfo:             vipInfo.GetBlogInfo(),
			ShortMessage:         vipInfo.GetShortMessage(),
			EndDate:              vipInfo.GetEndDate().Format(time.DateTime),
		}
		return nil
	})
}
