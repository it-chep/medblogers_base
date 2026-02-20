package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/change_vip_info/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
	"time"
)

func (i *Implementation) ChangeDoctorVipInfo(ctx context.Context, req *desc.ChangeDoctorVipInfoRequest) (resp *desc.ChangeDoctorVipInfoResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{doctor_id}/change_vip_info", func(ctx context.Context) error {
		updateReq, err := prepareUpdateVipInfoRequest(req)
		if err != nil {
			return err
		}

		err = i.admin.Actions.DoctorModule.DoctorAgg.ChangeVipInfo.Do(ctx, req.GetDoctorId(), updateReq)
		if err != nil {
			return err
		}

		return nil
	})
}

func prepareUpdateVipInfoRequest(req *desc.ChangeDoctorVipInfoRequest) (dto.UpdateRequest, error) {
	endDate, err := time.Parse(time.DateTime, req.GetEndDate())
	if err != nil {
		return dto.UpdateRequest{}, err
	}

	return dto.UpdateRequest{
		CanBarter:            req.GetCanBarter(),
		CanBuyAdvertising:    req.GetCanBuyAdvertising(),
		CanSellAdvertising:   req.GetCanSellAdvertising(),
		ShortMessage:         req.GetShortMessage(),
		BlogInfo:             req.GetBlogInfo(),
		AdvertisingPriceFrom: req.GetAdvertisingPriceFrom(),
		EndDate:              endDate,
	}, nil
}
