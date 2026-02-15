package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
)

// GetDoctorVip - /api/v1/doctors/{doctor_id}/vip [GET]
func (i *Implementation) GetDoctorVip(ctx context.Context, req *desc.GetDoctorVipRequest) (*desc.GetDoctorVipResponse, error) {
	vipInfo, err := i.doctors.Actions.DoctorVipInfo.Do(ctx, req.GetDoctorSlug())
	if err != nil {
		return nil, err
	}

	return &desc.GetDoctorVipResponse{
		CanBarter:            vipInfo.GetCanBarter(),
		CanBuyAdvertising:    vipInfo.GetCanBuyAdvertising(),
		CanSellAdvertising:   vipInfo.GetCanSellAdvertising(),
		AdvertisingPriceFrom: vipInfo.GetAdvertisingPriceFrom(),
		ShortMessage:         vipInfo.GetShortMessage(),
		BlogInfo:             vipInfo.GetBlogInfo(),
	}, nil
}
