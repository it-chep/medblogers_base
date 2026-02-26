package vip

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/domain/vip_card"
	"medblogers_base/internal/pkg/logger"
)

type Dal interface {
	GetVipInfo(ctx context.Context, doctorIDs []int64) (map[int64]*vip_card.VipCard, error)
}

type Service struct {
	dal Dal
}

func NewService(dal Dal) *Service {
	return &Service{
		dal: dal,
	}
}

// GetDoctorsVipInfo получение информации о випках
func (s *Service) GetDoctorsVipInfo(ctx context.Context, doctorIDs []int64) map[int64]dto.VipInfo {
	domainMap, err := s.dal.GetVipInfo(ctx, doctorIDs)
	if err != nil {
		logger.Error(ctx, "Ошибка получение випок %s", err)
		return nil
	}

	dtoMap := map[int64]dto.VipInfo{}
	for doctorID, vip := range domainMap {
		dtoMap[doctorID] = dto.VipInfo{
			CanBarter:            vip.GetCanBarter(),
			CanBuyAdvertising:    vip.GetCanBuyAdvertising(),
			CanSellAdvertising:   vip.GetCanSellAdvertising(),
			AdvertisingPriceFrom: vip.GetAdvertisingPriceFrom(),
		}
	}

	return dtoMap
}
