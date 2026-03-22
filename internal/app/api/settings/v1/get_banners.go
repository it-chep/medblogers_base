package v1

import (
	"context"

	"github.com/samber/lo"

	"medblogers_base/internal/modules/settings/action/get_banners/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/settings/v1"
)

func (i *Implementation) GetBanners(ctx context.Context, _ *desc.GetBannersRequest) (*desc.GetBannersResponse, error) {
	banners, err := i.settings.Actions.GetBanners.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetBannersResponse{
		Banners: lo.Map(banners, func(item dto.Banner, _ int) *desc.GetBannersResponse_Banner {
			return &desc.GetBannersResponse_Banner{
				Id:              item.ID,
				Name:            item.Name.String,
				OrderingNumber:  item.OrderingNumber,
				DesktopImage:    item.DesktopImage.String,
				DesktopFileType: item.DesktopFileType.String,
				MobileImage:     item.MobileImage.String,
				MobileFileType:  item.MobileFileType.String,
				BannerLink:      item.BannerLink.String,
			}
		}),
	}, nil
}
