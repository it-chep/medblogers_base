package v1

import (
	"fmt"

	adminDto "medblogers_base/internal/modules/admin/entities/banner/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

func (i *Implementation) newBannerResponse(item adminDto.Banner) *desc.Banner {
	return &desc.Banner{
		Id:              item.ID,
		Name:            item.Name,
		IsActive:        item.IsActive,
		OrderingNumber:  item.OrderingNumber,
		DesktopImage:    i.getImageURL(item.DesktopImage, item.DesktopFileType),
		DesktopFileType: item.DesktopFileType,
		MobileImage:     i.getImageURL(item.MobileImage, item.MobileFileType),
		MobileFileType:  item.MobileFileType,
		BannerLink:      item.BannerLink,
	}
}

func newUpdateRequest(reqName string, orderingNumber int64, bannerLink string) adminDto.UpdateRequest {
	return adminDto.UpdateRequest{
		Name:           reqName,
		OrderingNumber: orderingNumber,
		BannerLink:     bannerLink,
	}
}

func (i *Implementation) getImageURL(imageID, fileType string) string {
	if imageID == "" || fileType == "" || i.settingsBucket == "" {
		return ""
	}

	return fmt.Sprintf("https://storage.yandexcloud.net/%s/images/%s.%s", i.settingsBucket, imageID, fileType)
}
