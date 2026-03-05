package v1

import (
	"context"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	"medblogers_base/internal/pkg/formatters"
)

// GetDoctorVip - /api/v1/doctors/{doctor_id}/vip [GET]
func (i *Implementation) GetDoctorVip(ctx context.Context, req *desc.GetDoctorVipRequest) (*desc.GetDoctorVipResponse, error) {
	vipInfo, err := i.doctors.Actions.DoctorVipInfo.Do(ctx, req.GetDoctorSlug())
	if err != nil {
		return nil, err
	}

	if vipInfo == nil {
		return nil, status.Errorf(codes.NotFound, "doctor with slug '%s' not found or not VIP", req.GetDoctorSlug())
	}

	return &desc.GetDoctorVipResponse{
		CanBarter:            vipInfo.GetCanBarter(),
		CanBuyAdvertising:    vipInfo.GetCanBuyAdvertising(),
		CanSellAdvertising:   vipInfo.GetCanSellAdvertising(),
		AdvertisingPriceFrom: lo.Ternary(vipInfo.GetAdvertisingPriceFrom() > 0, formatters.HumanPrice(vipInfo.GetAdvertisingPriceFrom()), ""),
		ShortMessage:         vipInfo.GetShortMessage(),
		BlogInfo:             vipInfo.GetBlogInfo(),
	}, nil
}
