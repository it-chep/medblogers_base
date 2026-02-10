package converter

import (
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/client/subscribers/dto"
	"medblogers_base/internal/modules/admin/client/subscribers/indto"
)

func UpdateSubscribersRequestFromInDTO(items []indto.UpdateSubscribersItem) dto.UpdateSubscribersRequest {
	return dto.UpdateSubscribersRequest{
		Items: lo.Map(items, func(item indto.UpdateSubscribersItem, _ int) dto.SubscribersItem {
			return dto.SubscribersItem{
				Key:       item.Key.String(),
				SubsCount: item.SubsCount,
			}
		}),
	}
}
