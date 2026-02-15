package dto

import "github.com/samber/lo"

type Doctor struct {
	ID   int64
	Name string
	Slug string

	InstLink          string
	InstSubsCount     string
	InstSubsCountText string

	TgLink          string
	TgSubsCount     string
	TgSubsCountText string

	YouTubeLink          string
	YouTubeSubsCount     string
	YouTubeSubsCountText string

	VkLink          string
	VkSubsCount     string
	VkSubsCountText string

	Specialities []Speciality // Строка из основной и дополнительных специальностей
	Cities       []City       // Строка из основного и дополнительных городов

	Image string

	MainCityID       int64
	MainSpecialityID int64

	S3Key string

	IsKFDoctor bool
	IsVip      bool
}

type Doctors []Doctor

// GetVipIDs получение только ID випов
func (d Doctors) GetVipIDs() []int64 {
	return lo.FilterMap(d, func(item Doctor, index int) (int64, bool) {
		return item.ID, !item.IsVip
	})
}

type City struct {
	ID   int64
	Name string
}

type Speciality struct {
	ID   int64
	Name string
}

type VipInfo struct {
	CanBarter            bool
	CanBuyAdvertising    bool
	CanSellAdvertising   bool
	AdvertisingPriceFrom int64
}

type Response struct {
	Doctors []Doctor
	Vip     map[int64]VipInfo

	Pages            int64
	SubscribersCount string
}
